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
type TypeLeaveController struct{}

var typeLeaveModel = new(models.TypeLeaveModel)
var typeLeaveForm = new(forms.TypeLeaveForm)

// One...
func (ctrl TypeLeaveController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := typeLeaveModel.One(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "TypeLeave not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// All ...
func (ctrl TypeLeaveController) All(c *gin.Context) {
	results, err := typeLeaveModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get TypeLeave", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Create ...
func (ctrl TypeLeaveController) Create(c *gin.Context) {
	var form forms.TypeLeaveCreateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := typeLeaveForm.Create(validationErr)

		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	id, err := models.FlexibleInsert("sc_attendance.type_leave", form, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A TypeLeave with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "TypeLeave could not be created", nil)
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
		UserID: parsedUsrId, Type: "TypeLeave", Detail: "Created TypeLeave",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "TypeLeave created", gin.H{"id": id})
}

// Update ...
func (ctrl TypeLeaveController) Update(c *gin.Context) {
	id := c.Param("id")

	var form forms.TypeLeaveUpdateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := typeLeaveForm.Update(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_attendance.type_leave", form, cond, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A TypeLeave with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "TypeLeave could not be updated", nil)
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
		UserID: parsedUsrId, Type: "TypeLeave", Detail: "Updated TypeLeave",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "TypeLeave updated", nil)
}

// Delete ...
func (ctrl TypeLeaveController) Delete(c *gin.Context) {
	id := c.Param("id")

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleDelete("sc_attendance.type_leave", cond)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "TypeLeave could not be deleted", nil)
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
		UserID: parsedUsrId, Type: "TypeLeave", Detail: "Deleted TypeLeave",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "TypeLeave deleted", nil)
}
