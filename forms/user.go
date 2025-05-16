package forms

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Form ...
type UserForm struct{}

type UserCreateForm struct {
	Name           string    `form:"name" json:"name" binding:"required,min=3,max=100,fullName"`
	Email          string    `form:"email" json:"email" binding:"required,email"`
	Password       string    `form:"password" json:"password" binding:"required"`
	RoleId         uuid.UUID `form:"role_id" json:"role_id" binding:"required"`
	CompanyId 	   uuid.UUID `form:"company_id" json:"company_id" binding:"required"`
	PositionID     uuid.UUID `form:"position_id" json:"position_id" binding:"required"`
	BranchID       uuid.UUID `form:"branch_id" json:"branch_id" binding:"omitempty"`
	DepartmentId   uuid.UUID `form:"department_id" json:"department_id" binding:"omitempty"`
	Active         bool      `form:"active" json:"active" binding:"required"`
}

type CreateUserForm struct {
	Name           string    `form:"name" json:"name" binding:"required,min=3,max=100"`
	Email          string    `form:"email" json:"email" binding:"required,email"`
	Password       string    `form:"password" json:"password" binding:"required"`
	RoleId         uuid.UUID `form:"role_id" json:"role_id" binding:"required"`
}

type UserUpdateForm struct {
	Name           string    `form:"name" json:"name" binding:"omitempty,min=3,max=100,fullName"`
	Email          string    `form:"email" json:"email" binding:"omitempty,email"`
	Password       string    `form:"password" json:"password" binding:"omitempty"`
	RoleId         string    `form:"role_id" json:"role_id" binding:"omitempty"`
	CompanyId     uuid.UUID `form:"company_id" json:"company_id" binding:"omitempty"`
	PositionID     uuid.UUID `form:"position_id" json:"position_id" binding:"omitempty"`
	BranchID       uuid.UUID `form:"branch_id" json:"branch_id" binding:"omitempty"`
}

// LoginForm ...
type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

// Name ...
func (f UserForm) Name(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your name"
		}
		return errMsg[0]
	case "min", "max":
		return "Your name should be between 3 to 100 characters"
	case "fullName":
		return "Name should not include any special characters or numbers"
	default:
		return "Something went wrong, please try again later"
	}
}

// Email ...
func (f UserForm) Email(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your email"
		}
		return errMsg[0]
	case "min", "max", "email":
		return "Please enter a valid email"
	default:
		return "Something went wrong, please try again later"
	}
}

// Password ...
func (f UserForm) Password(tag string) (message string) {
	switch tag {
	case "required":
		return "Please enter your password"
	case "min", "max":
		return "Your password should be between 3 and 50 characters"
	case "eqfield":
		return "Your passwords does not match"
	default:
		return "Something went wrong, please try again later"
	}
}

// Signin ...
func (f UserForm) Login(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}
			if err.Field() == "Password" {
				return f.Password(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// User Create ...
func (f UserForm) UserCreate(err error) string {
	fmt.Println(err)
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// User Update ...
func (f UserForm) UserUpdate(err error) string {
	fmt.Println(err)
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return f.Name(err.Tag())
			}

			if err.Field() == "Email" {
				return f.Email(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
