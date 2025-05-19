package controllers

import (
	"fmt"
	"net/http"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/lib/pq"
)

// Controller ...
type LeaveController struct{}

var leaveModel = new(models.LeaveModel)
var leaveForm = new(forms.LeaveForm)

// AllByApprover
func (ctrl LeaveController) AllByApprover(c *gin.Context) {
	id := c.Param("id")

	data, err := leaveModel.AllByApprover(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Leave not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// OnBy DepartmentId
func (ctrl LeaveController) OneByDepartment(c *gin.Context) {
	id := c.Param("id")

	data, err := leaveModel.OneByDepartment(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Leave not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// All ...
func (ctrl LeaveController) All(c *gin.Context) {
	results, err := leaveModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Leave", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Pengajuan ...
func (ctrl LeaveController) Pengajuan(c *gin.Context) {
	var form forms.PengajuanForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := leaveForm.Pengajuan(validationErr)

		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	id, err := models.FlexibleInsert("sc_attendance.leave", form, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A leave with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Leave could not be created", nil)
		return
	}

	// parsing user id to UUID type
	parsedUsrId, parseErr := uuid.Parse(getUserID(c))
	if parseErr != nil {
		standarizedResponse(c, true, http.StatusBadRequest, "Failed to parsing uuid", nil)
		return
	}
	// logging activity
	actLog := &interType.LogActivity{
		UserID: parsedUsrId, Type: "leave", Detail: "Pengajuan Leave",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Pengajuan Leave Success", gin.H{"id": id})
}
