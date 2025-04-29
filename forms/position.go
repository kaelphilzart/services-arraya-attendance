package forms

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Form ...
type PositionForm struct{}

// Create ...
type PositionCreateForm struct {
	Name         string `form:"name" json:"name" binding:"required,min=3,max=100"`
	DepartmentId string `form:"department_id" json:"department_id" binding:"required"`
	Level        string `form:"level" json:"level" binding:"required"`
}

// Update ...
type PositionUpdateForm struct {
	Name         string `form:"name" json:"name" binding:"omitempty,min=3,max=100"`
	DepartmentId string `form:"department_id" json:"department_id" binding:"omitempty"`
	Level        string `form:"level" json:"level" binding:"omitempty"`
}

// Name ...
func (f PositionForm) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the name"
		}
		return errMsg[0]
	case "min", "max":
		return "Name should be between 3 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

// Create ...
func (f PositionForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f PositionForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		fmt.Println(err)

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
