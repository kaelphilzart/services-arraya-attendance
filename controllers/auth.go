package controllers

import (
	"fmt"
	"net/http"
	"os"

	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/google/uuid"
)

// Controller ...
type AuthController struct{}

var authModel = new(models.AuthModel)

// TokenResetPasswordValid ...
func (ctl AuthController) TokenResetPasswordValid(c *gin.Context) {
	tokenAuth, err := authModel.ExtractTokenResetPasswordMetadata(c.Request)
	if err != nil {
		//Token either expired or not valid
		standarizedResponse(c, true, http.StatusUnauthorized, "Please login first", nil)
		return
	}

	userID := tokenAuth.UserID

	//To be called from GetUserID()
	c.Set("userID", userID)
}

// TokenAdminValid ...
func (ctl AuthController) TokenAdminValid(c *gin.Context) {
	tokenAuth, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		fmt.Println("Token error:", err)
		standarizedResponse(c, true, http.StatusUnauthorized, "Please login first", nil)
		return
	}
	userID, err := authModel.FetchAuth(tokenAuth)
	if err != nil {
		fmt.Println("FetchAuth error:", err)
		standarizedResponse(c, true, http.StatusUnauthorized, "Please login first", nil)
		return
	}

	// fmt.Println("role_client from token:", tokenAuth.RoleClient) // <=== INI DIA

	if tokenAuth.RoleClient != "adm" {
		standarizedResponse(c, true, http.StatusForbidden, "Forbidden Access!!", nil)
		return
	}

	c.Set("userID", userID)
	c.Set("role", tokenAuth.RoleClient)
}

// TokenValid ...
func (ctl AuthController) TokenValid(c *gin.Context) {

	tokenAuth, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		//Token either expired or not valid
		standarizedResponse(c, true, http.StatusUnauthorized, "Please login first", nil)
		return
	}

	userID, err := authModel.FetchAuth(tokenAuth)
	if err != nil {
		//Token does not exists in Redis (User logged out or expired)
		standarizedResponse(c, true, http.StatusUnauthorized, "Please login first", nil)
		return
	}

	//To be called from GetUserID()
	c.Set("userID", userID)
	//To be called from GetRole()
	c.Set("role", tokenAuth.RoleClient)
}

// Refresh ...
func (ctl AuthController) Refresh(c *gin.Context) {
	var tokenForm forms.Token

	if c.ShouldBindJSON(&tokenForm) != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "Invalid form", gin.H{"form": tokenForm})
		return
	}

	//verify the token
	token, err := jwt.Parse(tokenForm.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			standarizedResponse(c, true, http.StatusUnauthorized, "unexpected signing method", nil)
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string) //convert the interface to string

		if !ok {
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}
		roleClient, ok := claims["role_client"].(string) //convert the interface to string

		if !ok {
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}

		userID := claims["user_id"].(string)

		authD := &interType.AccessDetails{
			AccessUUID: refreshUUID,
			UserID:     userID,
		}
		userID, err := authModel.FetchAuth(authD)

		if err != nil {
			//Token does not exists in Redis (User logged out or expired)
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}

		parsedID, err := uuid.Parse(userID)

		if err != nil {
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := authModel.DeleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 { //if any goes wrong
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}

		//Create new pairs of refresh and access tokens
		ts, createErr := authModel.CreateToken(parsedID, roleClient)
		if createErr != nil {
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}
		//save the tokens metadata to redis
		saveErr := authModel.CreateAuth(parsedID, ts)
		if saveErr != nil {
			standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		standarizedResponse(c, false, http.StatusOK, "Success", tokens)
	} else {
		standarizedResponse(c, true, http.StatusUnauthorized, "Invalid authorization, please login again", nil)
	}
}
