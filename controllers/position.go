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
type PositionController struct{}

var positionModel = new(models.PositionModel)
var positionForm = new(forms.PositionForm)

// One...
func (ctrl PositionController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := positionModel.One(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Position not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// All ...
func (ctrl PositionController) All(c *gin.Context) {
	results, err := positionModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Position", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Create ...
func (ctrl PositionController) Create(c *gin.Context) {
	var form forms.PositionCreateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := positionForm.Create(validationErr)

		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	id, err := models.FlexibleInsert("sc_user.position", form, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A position with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Position could not be created", nil)
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
		UserID: parsedUsrId, Type: "Position", Detail: "Created Position",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Position created", gin.H{"id": id})
}

// Update ...
func (ctrl PositionController) Update(c *gin.Context) {
	id := c.Param("id")

	var form forms.PositionUpdateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := positionForm.Update(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_user.position", form, cond, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A position with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Position could not be updated", nil)
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
		UserID: parsedUsrId, Type: "Position", Detail: "Updated Position",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Position updated", nil)
}

// Delete ...
func (ctrl PositionController) Delete(c *gin.Context) {
	id := c.Param("id")

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleDelete("sc_user.position", cond)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "Position could not be deleted", nil)
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
		UserID: parsedUsrId, Type: "Position", Detail: "Deleted Position",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Position deleted", nil)
}
