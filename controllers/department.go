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
type DepartmentController struct{}

var departmentModel = new(models.DepartmentModel)
var departmentForm = new(forms.DepartmentForm)

// One...
func (ctrl DepartmentController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := departmentModel.One(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Department not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// All ...
func (ctrl DepartmentController) All(c *gin.Context) {
	results, err := departmentModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Department", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Create ...
func (ctrl DepartmentController) Create(c *gin.Context) {
	var form forms.DepartmentCreateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := departmentForm.Create(validationErr)

		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	id, err := models.FlexibleInsert("sc_user.department", form, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A department with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Department could not be created", nil)
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
		UserID: parsedUsrId, Type: "Department", Detail: "Created Department",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Department created", gin.H{"id": id})
}

// Update ...
func (ctrl DepartmentController) Update(c *gin.Context) {
	id := c.Param("id")

	var form forms.DepartmentUpdateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := departmentForm.Update(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_user.department", form, cond, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A department with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Department could not be updated", nil)
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
		UserID: parsedUsrId, Type: "Department", Detail: "Updated Department",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Department updated", nil)
}

// Delete ...
func (ctrl DepartmentController) Delete(c *gin.Context) {
	id := c.Param("id")

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleDelete("sc_user.department", cond)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "Department could not be deleted", nil)
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
		UserID: parsedUsrId, Type: "Department", Detail: "Deleted Department",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Department deleted", nil)
}
   