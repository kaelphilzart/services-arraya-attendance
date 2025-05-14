package controllers

import (
	"net/http"
	"fmt"
	"services-arraya-attendance/forms"
	interType "services-arraya-attendance/interfaces"
	"services-arraya-attendance/models"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/lib/pq"
)

type ShiftController struct{}

var shiftModel = new(models.ShiftModel)
var shiftForm = new(forms.ShiftForm)

// One
func (ctrl ShiftController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := shiftModel.One(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Shift not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// All ...
func (ctrl ShiftController) All(c *gin.Context) {
	results, err := shiftModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Shift", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Create ...
func (ctrl ShiftController) Create(c *gin.Context) {
	var form forms.ShiftCreateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := shiftForm.Create(validationErr)

		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	id, err := models.FlexibleInsert("sc_attendance.Shift", form, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A Shift with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Shift could not be created", nil)
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
		UserID: parsedUsrId, Type: "Shift", Detail: "Created Shift",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Shift created", gin.H{"id": id})
}

// Update ...
func (ctrl ShiftController) Update(c *gin.Context) {
	id := c.Param("id")

	var form forms.ShiftUpdateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := shiftForm.Update(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_attendance.shift", form, cond, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A Shift with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Shift could not be updated", nil)
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
		UserID: parsedUsrId, Type: "Shift", Detail: "Updated Shift",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Shift updated", nil)
}

// Delete ...
func (ctrl ShiftController) Delete(c *gin.Context) {
	id := c.Param("id")

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleDelete("sc_attendance.shift", cond)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "Shift could not be deleted", nil)
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
		UserID: parsedUsrId, Type: "Shift", Detail: "Deleted Shift",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Shift deleted", nil)
}