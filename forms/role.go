package forms

import(
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

//form

type RoleForm struct{}

type RoleCreateForm struct {
	Name           string    `form:"name" json:"name" binding:"required,min=2,max=50"`
	SlugName       string    `form:"slug_name" json:"slug_name" binding:"required"`
}

type UpdateRoleForm struct {
	Name        string `form:"name" json:"name" binding:"omitempty,min=3,max=255,fullName"`
	SlugName    string `form:"slug_name" json:"slug_name" binding:"omitempty"`
}



// Slugname..
func (f RoleForm) SlugName(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter your slugname"
		}
		return errMsg[0]
	case "min", "max", "slugname":
		return "Please enter a valid slugname"
	default:
		return "Something went wrong, please try again later"
	}
}

//create
func (f RoleForm) Create(err error) string {
	fmt.Println(err)
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Slugname" {
				return f.SlugName(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

//update
func (f RoleForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Slugname" {
				return f.SlugName(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
