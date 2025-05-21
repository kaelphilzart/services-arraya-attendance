package forms

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Form ...
type BranchForm struct{}

// Create ...
type BranchCreateForm struct {
	CompanyId	   string `form:"company_id" json:"company_id" binding:"required"`
	Name           string `form:"name" json:"name" binding:"required,min=3"`
	Address        string `form:"address" json:"address" binding:"required"`
	Contact    	   string `form:"contact" json:"contact" binding:"omitempty"`
}

// Update ...
type BranchUpdateForm struct {
	CompanyId string `form:"company_id" json:"company_id" binding:"omitempty"`
	Name           string `form:"name" json:"name" binding:"omitempty,min=3"`
	Address        string `form:"address" json:"address" binding:"omitempty"`
	Contact    	   string `form:"contact" json:"contact" binding:"omitempty"`
}

// Name ...
func (f BranchForm) Name(tag string, errMsg ...string) (message string) {
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
func (f BranchForm) Create(err error) string {
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
func (f BranchForm) Update(err error) string {
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
