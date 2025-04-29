package controllers

import (
	"fmt"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Controller ...
type UserController struct{}

var userModel = new(models.UserModel)
var logActivityModel = new(models.LogActivityModel)
var userForm = new(forms.UserForm)

// Gabungan form user + profile
type UserWithProfileForm struct {
	forms.UserCreateForm
	forms.UserProfileCreateForm
}

// getUserID ...
func getUserID(c *gin.Context) (userID string) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(string)
}

// Login ...
func (ctrl UserController) Login(c *gin.Context) {
    var loginForm forms.LoginForm

    if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
        message := userForm.Login(validationErr)
        standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
        return
    }

    // Call the Login function to get user and token
    user, token, err := userModel.Login(loginForm)
    if err != nil {
        fmt.Println("Login error:", err)
        standarizedResponse(c, true, http.StatusUnauthorized, "Invalid email or password. Please try again", nil)
        return
    }

    // Debugging the token to see if it's empty or valid
    fmt.Println("Access Token:", token.AccessToken)
    fmt.Println("Refresh Token:", token.RefreshToken)

    // Proceed with logging activity
    parsedUsrId := user.ID
    actLog := &interType.LogActivity{
        UserID: parsedUsrId, Type: "Auth", Detail: "User logged in",
    }
    go models.LogActivity(actLog)

    // Respond back to the client
    standarizedResponse(c, false, http.StatusOK, "Successfully login", gin.H{
        "user":  user,
        "token": token,
    })
}


// Logout ...
func (ctrl UserController) Logout(c *gin.Context) {
	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "User not logged in", nil)
		return
	}

	deleted, delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil || deleted == 0 { //if any goes wrong
		standarizedResponse(c, true, http.StatusUnauthorized, "Invalid request", nil)
		return
	}

	// parsing user id to UUID type
	parsedUsrId, parseErr := uuid.Parse(au.UserID)
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
		return
	}
	// logging activity
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "Auth", Detail: "User logged out",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Successfully logged out", nil)
}

// LogActivityAll ...
func (ctrl UserController) LogActivityAll(c *gin.Context) {
	id := getUserID(c)

	ParamsLogActivity := &interType.ParamsLogActivityAll{
		UserID: id,
	}

	results, err := logActivityModel.All(ParamsLogActivity)
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Log Activity", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// All ...
func (ctrl UserController) All(c *gin.Context) {

	results, err := userModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Users", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// One ...
func (ctrl UserController) One(c *gin.Context) {

	id := c.Param("id")

	data, err := userModel.One(id)
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get User", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// Create ...
func (ctrl UserController) Create(c *gin.Context) {
	var form UserWithProfileForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := userForm.UserCreate(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to hash password", nil)
		return
	}

	form.Password = string(hashedPassword)

	// Insert ke tabel user
	userData := forms.UserCreateForm{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
		RoleId:   form.RoleId,
		// tambahkan field lain jika perlu
	}

	userID, err := models.FlexibleInsert("sc_users.user", &userData, "id")
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "tb_user_unique_email"`) {
			standarizedResponse(c, true, http.StatusBadRequest, "Email has already been registered", nil)
		} else {
			standarizedResponse(c, true, http.StatusBadRequest, "User could not be created", nil)
		}
		return
	}

	// Insert ke tabel user_profile
	userProfileData := forms.UserProfileCreateForm{
		UserId:      userID,
		FullName:    form.FullName,
		BirthDate:   form.BirthDate,
		BirthPlace:  form.BirthPlace,
		PhoneNumber: form.PhoneNumber,
		Address:     form.Address,
		Gender:      form.Gender,
		Photo:       form.Photo,
	}

	_, err = models.FlexibleInsert("sc_users.user_profile", &userProfileData, "id")
	if err != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "User created but profile failed", nil)
		return
	}

	// Logging activity
	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
	if parseErr == nil {
		actLog := &interType.LogActivity{
			UserID: parsedUsrId,
			Type:   "User Management",
			Detail: "Created User + Profile",
		}
		go models.LogActivity(actLog)
	}

	standarizedResponse(c, false, http.StatusOK, "Successfully created user & profile", gin.H{"id": userID})
}

// Update ...
func (ctrl UserController) Update(c *gin.Context) {
	id := c.Param("id")

	var form forms.UserUpdateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := userForm.UserUpdate(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}
	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_users.users", form, cond, "id")
	if err != nil {
		fmt.Println("ini error", err.Error())

		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_unique_email"`) {
			standarizedResponse(c, true, http.StatusBadRequest, "Email has already been registered", nil)
		} else {
			standarizedResponse(c, true, http.StatusBadRequest, "User could not be created", nil)
		}

		return
	}

	// ------------------------------
	// parsing user id to UUID type
	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
		return
	}
	// logging activity
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "User Management", Detail: "Updated User",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Successfully Update User", nil)
}

// Delete ...
func (ctrl UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleDelete("sc_users.users", cond)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "User could not be deleted", nil)
		return
	}

	// ------------------------------
	// parsing user id to UUID type
	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
		return
	}
	// logging activity
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "User Management", Detail: "Deleted User",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Successfully Delete User", nil)
}


// func (ctrl UserController) CreateUser(c *gin.Context) {
// 	var form forms.CreateUserForm

// 	// Validasi input
// 	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
// 		message := userForm.UserCreate(validationErr)
// 		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
// 		return
// 	}

// 	// Hash password
// 	bytePassword := []byte(form.Password)
// 	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	if err != nil {
// 		standarizedResponse(c, true, http.StatusBadRequest, "Failed to hashing new password", nil)
// 		return
// 	}

// 	// Membuat form baru dengan password yang sudah di-hash
// 	formNewWithPass := &forms.UserCreateForm{}
// 	mergeStructs(form, formNewWithPass)
// 	formNewWithPass.Password = string(hashedPassword)

// 	// Insert ke database
// 	id, err := models.FlexibleInsert("sc_users.users", formNewWithPass, "id")
// 	if err != nil {
// 		fmt.Println("ini error", err.Error())

// 		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "tb_user_unique_nip"`) {
// 			standarizedResponse(c, true, http.StatusBadRequest, "NIP has already been registered", nil)
// 		} else if strings.Contains(err.Error(), `duplicate key value violates unique constraint "tb_user_unique_email"`) {
// 			standarizedResponse(c, true, http.StatusBadRequest, "Email has already been registered", nil)
// 		} else {
// 			standarizedResponse(c, true, http.StatusBadRequest, "User could not be created", nil)
// 		}
// 		return
// 	}

// 	// ------------------------------
// 	// DISINI kita skip log activity karena belum ada user ID yang login
// 	// (atau kalau mau tetap logging, pakai ID kosong / "SYSTEM" user)
// 	// ------------------------------

// 	// Langsung kirim response sukses
// 	standarizedResponse(c, false, http.StatusOK, "Successfully Create User", gin.H{"id": id})
// }

