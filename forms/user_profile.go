package forms

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Form ...
type UserProfileForm struct{}

type UserProfileCreateForm struct {
	UserId      uuid.UUID `form:"user_id" json:"user_id" binding:"required"`
	FullName    string `form:"full_name" json:"full_name" binding:"required,min=5,max=100,fullName"`
	BirthDate   string `form:"birth_date" json:"birth_date" binding:"required"`
	BirthPlace  string `form:"birth_place" json:"birth_place" binding:"required"`
	Address     string `form:"address" json:"address" binding:"required"`
	Gender      string `form:"gender" json:"gender" binding:"required"`
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required,min=12,max=14"`
	Photo    	string `form:"photo" json:"photo" binding:"omitempty"`
}

type UpdateUserProfileForm struct {
	FullName    string `form:"full_name" json:"full_name" binding:"omitempty,min=3,max=100,fullName"`
	BirthDate   string `form:"birth_date" json:"birth_date" binding:"omitempty"`
	BirthPlace  string `form:"birth_place" json:"birth_place" binding:"omitempty"`
	Address     string `form:"address" json:"address" binding:"omitempty"`
	Gender      string `form:"gender" json:"gender" binding:"omitempty"`
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"omitempty,min=10,max=14"`
	Photo       string `form:"photo" json:"photo" binding:"omitempty"`
}

func (f UserProfileForm) UpdateUserProfile(err error) string {
	fmt.Println(err)
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "PhoneNumber" {
				return f.PhoneNumber(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// PhoneNumber ...
func (f UserProfileForm) PhoneNumber(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your phone number"
		}
		return errMsg[0]
	case "min", "max":
		return "Phone number should be between 12 to 14 digits"
	case "numeric":
		return "Phone number must contain only digits"
	default:
		return "Something went wrong, please try again later"
	}
}