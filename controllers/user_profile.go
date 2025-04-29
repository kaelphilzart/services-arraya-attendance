package controllers

import (
	"fmt"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileController struct{}

var userProfileModel = new(models.UserProfileModel)
var UpdateuserProfileForm = new(forms.UserProfileForm)

// Update profile...
func (ctrl UserProfileController) Update(c *gin.Context) {
	id := getUserID(c)

	var form forms.UpdateUserProfileForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := UpdateuserProfileForm.UpdateUserProfile(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}
	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_user.user_profile", form, cond, "id")
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Profile could not be updated", nil)
		return
	}
	// parsing user id to UUID type
	parsedUsrId, parseErr := uuid.Parse(id)
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
		return
	}
	// logging activity
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "User", Detail: "Updated Profile",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Successfully Update Profile", nil)
}

// one
func (ctrl UserProfileController) One(c *gin.Context) {

	id := c.Param("id")

	data, err := userProfileModel.One(id)
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get User", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}