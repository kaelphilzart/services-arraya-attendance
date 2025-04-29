package models

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"services-arraya-attendance/db"
	interType "services-arraya-attendance/interfaces"

	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/google/uuid"
)

// AuthModel ...
type AuthModel struct{}

// CreateToken ...
func (m AuthModel) CreateToken(userID uuid.UUID, roleClient string) (*interType.TokenDetails, error) {

    td := &interType.TokenDetails{}
    tokenUUID := uuid.New().String()
    
    td.AtExpires = time.Now().Add(time.Minute * 600).Unix()  // Access token expiry
    td.AccessUUID = tokenUUID

    td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()  // Refresh token expiry
    td.RefreshUUID = tokenUUID

    var err error
    // Creating Access Token
    atClaims := jwt.MapClaims{}
    atClaims["authorized"] = true
    atClaims["access_uuid"] = td.AccessUUID
    atClaims["user_id"] = userID
    atClaims["role_client"] = roleClient
    atClaims["exp"] = td.AtExpires

    at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
    td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
    if err != nil {
        return nil, err
    }

    // Creating Refresh Token
    rtClaims := jwt.MapClaims{}
    rtClaims["refresh_uuid"] = td.RefreshUUID
    rtClaims["user_id"] = userID
    rtClaims["role_client"] = roleClient
    rtClaims["exp"] = td.RtExpires
    rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
    td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
    if err != nil {
        return nil, err
    }

    return td, nil
}

// CreateTokenResetPassword ...
func (m AuthModel) CreateTokenResetPassword(userID uuid.UUID) (*interType.TokenDetails, error) {

	td := &interType.TokenDetails{}
	tokenUUID := uuid.New().String()
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = tokenUUID

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = tokenUUID

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("RESET_PASS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// CreateAuth ...
func (m AuthModel) CreateAuth(userid uuid.UUID, td *interType.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	errAccess := db.GetRedis().Set(td.AccessUUID, userid.String(), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := db.GetRedis().Set(td.RefreshUUID, userid.String(), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// ExtractToken ...
func (m AuthModel) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyTokenResetPassword ...
func (m AuthModel) VerifyTokenResetPassword(r *http.Request) (*jwt.Token, error) {
	tokenString := m.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("RESET_PASS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// VerifyToken ...
func (m AuthModel) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := m.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
func (m AuthModel) TokenValid(r *http.Request) error {
	token, err := m.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenResetPasswordMetadata ...
func (m AuthModel) ExtractTokenResetPasswordMetadata(r *http.Request) (*interType.AccessDetails, error) {
	token, err := m.VerifyTokenResetPassword(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID := claims["user_id"].(string)
		return &interType.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

// ExtractTokenMetadata ...
func (m AuthModel) ExtractTokenMetadata(r *http.Request) (*interType.AccessDetails, error) {
	token, err := m.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID := claims["user_id"].(string)
		roleClient, ok := claims["role_client"].(string)
		if !ok {
			return nil, err
		}
		return &interType.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
			RoleClient: roleClient,
		}, nil
	}
	return nil, err
}

// FetchAuth ...
func (m AuthModel) FetchAuth(authD *interType.AccessDetails) (string, error) {
	userid, err := db.GetRedis().Get(authD.AccessUUID).Result()
	if err != nil {
		return "0", err
	}
	userID := userid
	return userID, nil
}

// DeleteAuth ...
func (m AuthModel) DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := db.GetRedis().Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
