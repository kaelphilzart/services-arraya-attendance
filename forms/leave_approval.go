package forms

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

type LeaveApprovalForm struct{}

// Approval
type ApproveForm struct {
	LeaveId                  string    `form:"leave_id" json:"leave_id" binding:"required"`
	Level               	 string    `form:"level" json:"level" binding:"omitempty"`
	ApprovedBy  			 string    `form:"approved_by" json:"approved_by" binding:"required"`
	Status          		 bool       `form:"status" json:"status" binding:"required"`
	Note					 string    `form:"note" json:"note" binding:"required"`
	ApprovedAt	 			 string    `form:"approved_at" json:"approved_at" binding:"required"`
}


func (f LeaveApprovalForm) Approval(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Terjadi kesalahan pada format data, silakan coba lagi."
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			switch fieldErr.Field() {
			case "Note":
				return "Note wajib diisi."
			}
		}

	default:
		return "Permintaan tidak valid, silakan cek kembali data Anda."
	}

	return "Terjadi kesalahan, silakan coba lagi nanti."
}
