package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// Form ...
type LeaveForm struct{}

// Pengajuan
type PengajuanForm struct {
	UserID                   string    `form:"user_id" json:"user_id" binding:"omitempty"`
	TypeLeaveId 			 string    `form:"type_leave_id" json:"type_leave_id" binding:"required"`
	UrlPhoto               	 string    `form:"url_photo" json:"url_photo" binding:"omitempty"`
	StartDate  			 	 string    `form:"start_date" json:"start_date" binding:"required"`
	EndDate 			 	 string    `form:"end_date" json:"end_date" binding:"required"`
	Status          		 bool      `form:"status" json:"status" binding:"required"`
	CurrentApprovalLevel	 int8      `form:"current_approval_level" json:"current_approval_level" binding:"required"`
	Description	 			 string    `form:"description" json:"description" binding:"omitempty"`
	
}

// Update Pengajuan
type UpdatePengajuanForm struct {
	Status 			 		 string    `form:"status" json:"status" binding:"omitempty"`
	CurrentApprovalLevel	 int8    `form:"current_approval_level" json:"current_approval_level" binding:"omitempty"`
}


func (f LeaveForm) Pengajuan(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Terjadi kesalahan pada format data, silakan coba lagi."
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			switch fieldErr.Field() {
			case "TypeLeaveId":
				return "Tipe cuti/izin wajib diisi."
			case "StartDate":
				return "Tanggal mulai cuti wajib diisi."
			case "EndDate":
				return "Tanggal selesai cuti wajib diisi."
			case "Status":
				return "Status pengajuan wajib diisi."
			case "CurrentApprovalLevel":
				return "Level approval saat ini wajib diisi."
			}
		}

	default:
		return "Permintaan tidak valid, silakan cek kembali data Anda."
	}

	return "Terjadi kesalahan, silakan coba lagi nanti."
}