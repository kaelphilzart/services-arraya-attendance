package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// Form ...
type ShiftForm struct{}

// Shift Create ...
type ShiftCreateForm struct {
	CompanyId string `form:"company_id" json:"company_id" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required,min=3,max=255"`
	StartTime string `form:"start_time" json:"start_time" binding:"required"`
	EndTime   string `form:"end_time" json:"end_time" binding:"required"`
}

// Shift Update ...
type ShiftUpdateForm struct {
	CompanyId string `form:"company_id" json:"company_id" binding:"omitempty"`
	Name      string `form:"name" json:"name" binding:"omitempty,min=3,max=255"`
	StartTime string `form:"start_time" json:"start_time" binding:"omitempty"`
	EndTime   string `form:"end_time" json:"end_time" binding:"omitempty"`
}

// Create ...
func (f ShiftForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Update ...
func (f ShiftForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
