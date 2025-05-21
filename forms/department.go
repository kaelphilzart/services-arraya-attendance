package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// Form ...
type DepartmentForm struct{}

// Create ...
type DepartmentCreateForm struct {
	Name            string `form:"name" json:"name" binding:"required,min=3"`
	CompanyId       string `form:"company_id" json:"company_id" binding:"omitempty"`
	BranchId        string `form:"branch_id" json:"branch_id" binding:"omitempty"`
	DirectorId      string `form:"director_id" json:"director_id" binding:"omitempty"`
}

// Update ...
type DepartmentUpdateForm struct {
	Name            string `form:"name" json:"name" binding:"required,min=3"`
	CompanyId       string `form:"company_id" json:"company_id" binding:"omitempty"`
	BranchId        string `form:"branch_id" json:"branch_id" binding:"omitempty"`
	DirectorId      string `form:"director_id" json:"director_id" binding:"omitempty"`
}

// Name ...
func (f DepartmentForm) Name(tag string, errMsg ...string) (message string) {
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
func (f DepartmentForm) Create(err error) string {
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
func (f DepartmentForm) Update(err error) string {
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
