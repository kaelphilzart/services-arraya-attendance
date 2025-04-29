package forms

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

// Form
type AttendanceForm struct{}

// Attendance In
type AttendanceInForm struct {
	UserID                   string    `form:"user_id" json:"user_id" binding:"omitempty"`
	Date                     time.Time `form:"date" json:"date" binding:"omitempty"`
	ChekInTime               string    `form:"chek_in_time" json:"chek_in_time" binding:"required"`
	LatitudeIn  			 string    `form:"latitude_in" json:"latitude_in" binding:"required"`
	LongitudeIn 			 string    `form:"longitude_in" json:"longitude_in" binding:"required"`
	PhotoIn          		 string    `form:"photo_in" json:"photo_in" binding:"required"`
}

// Attendance Out
type AttendanceOutForm struct {
	UserID                   string    `form:"user_id" json:"user_id" binding:"omitempty"`
	Date                     time.Time `form:"date" json:"date" binding:"omitempty"`
	ChekOutTime              string    `form:"chek_out_time" json:"chek_out_time" binding:"required"`
	LatitudeOut  			 string    `form:"latitude_out" json:"latitude_out" binding:"required"`
	LongitudeOut 			 string    `form:"longitude_out" json:"longitude_out" binding:"required"`
	PhotoOut          		 string    `form:"photo_out" json:"photo_out" binding:"required"`
}

// Attendance In ...
func (f AttendanceForm) AttendanceIn(err error) string {
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

// Attendance Out ...
func (f AttendanceForm) AttendanceOut(err error) string {
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