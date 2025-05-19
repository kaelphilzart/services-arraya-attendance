package controllers

import (
	"fmt"
	"net/http"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/lib/pq"
)

// Controller ...
type LeaveApprovalController struct{}

var leaveApprovalModel = new(models.LeaveApprovalModel)
var leaveApprovalForm = new(forms.LeaveApprovalForm)



// OneByLeaveId
func (ctrl LeaveApprovalController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := leaveApprovalModel.One(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Approval not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}


// All ...
func (ctrl LeaveApprovalController) All(c *gin.Context) {
	results, err := leaveApprovalModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Leave", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Approval ...
func (ctrl LeaveApprovalController) Approval(c *gin.Context) {
	var form forms.ApproveForm
	var updateForm forms.UpdatePengajuanForm

	// Bind JSON
	if err := c.ShouldBindJSON(&form); err != nil {
		message := leaveApprovalForm.Approval(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	// Insert approval ke leave_approval
	id, err := models.FlexibleInsert("sc_attendance.leave_approval", form, "id")
	if err != nil {
		fmt.Println(err)
		if pqErr, ok := err.(*pq.Error); ok && strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
			standarizedResponse(c, true, http.StatusConflict, "An approval with this name already exists", nil)
			return
		}
		standarizedResponse(c, true, http.StatusNotAcceptable, "Approval could not be created", nil)
		return
	}

	// Ambil data cuti saat ini
	leaveData, err := leaveModel.GetStatus(form.LeaveId)
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotFound, "Leave not found", nil)
		return
	}

	// Tambah level jika status == true
	newCurrentLevel := leaveData.CurrentApprovalLevel
	if form.Status {
		newCurrentLevel += 1
	}

	// Set nilai update form
	updateForm.CurrentApprovalLevel = int8(newCurrentLevel)
	updateForm.Status = strconv.FormatBool(form.Status) // jadi string "true" atau "false"

	// Update leave
	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: form.LeaveId,
	}

	updateErr := models.FlexibleUpdate("sc_attendance.leave", updateForm, cond, "id")
	if updateErr != nil {
		fmt.Println(updateErr)
		if pqErr, ok := updateErr.(*pq.Error); ok && strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
			standarizedResponse(c, true, http.StatusConflict, "A department with this name already exists", nil)
			return
		}
		standarizedResponse(c, true, http.StatusNotAcceptable, "Leave could not be updated", nil)
		return
	}

	// Parse UUID user
	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parse UUID", nil)
		return
	}

	// Log aktivitas
	actLog := &interType.LogActivity{
		UserID: parsedUsrId,
		Type:   "approve",
		Detail: "Approval Leave",
	}
	go models.LogActivity(actLog)

	// Berhasil
	standarizedResponse(c, false, http.StatusOK, "Approval Leave Success", gin.H{"id": id})
}



