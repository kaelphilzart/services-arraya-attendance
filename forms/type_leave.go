package forms

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Form ...
type TypeLeaveForm struct{}

// Create ...
type TypeLeaveCreateForm struct {
	Code         string `form:"code" json:"code" binding:"required,min=2,max=100"`
	Name         string `form:"name" json:"name" binding:"required,min=3,max=100"`
}

// Update ...
type TypeLeaveUpdateForm struct {
	Code         string `form:"code" json:"code" binding:"omitempty,min=2,max=100"`
	Name         string `form:"name" json:"name" binding:"omitempty,min=3,max=100"`
}

// Name ...
func (f TypeLeaveForm) Code(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the code"
		}
		return errMsg[0]
	case "min", "max":
		return "Code should be between 2 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Create ...
func (f TypeLeaveForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Code" {
				return f.Code(err.Tag())
			}

		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f TypeLeaveForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		fmt.Println(err)

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Code" {
				return f.Code(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
