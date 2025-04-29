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
type BranchController struct{}

var branchModel = new(models.BranchModel)
var branchForm = new(forms.BranchForm)

// One...
func (ctrl BranchController) One(c *gin.Context) {
	id := c.Param("id")

	data, err := branchModel.One(id)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotFound, "Branch not found", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", data)
}

// All ...
func (ctrl BranchController) All(c *gin.Context) {
	results, err := branchModel.All()
	if err != nil {
		fmt.Println(err)
		standarizedResponse(c, true, http.StatusNotAcceptable, "Could not get Branch", nil)
		return
	}

	standarizedResponse(c, false, http.StatusOK, "Success", results)
}

// Create ...
func (ctrl BranchController) Create(c *gin.Context) {
	var form forms.BranchCreateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := branchForm.Create(validationErr)

		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	id, err := models.FlexibleInsert("sc_user.branch", form, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A branch with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Branch could not be created", nil)
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
		UserID: parsedUsrId, Type: "Branch", Detail: "Created Branch",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Branch created", gin.H{"id": id})
}

// Update ...
func (ctrl BranchController) Update(c *gin.Context) {
	id := c.Param("id")

	var form forms.BranchUpdateForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := branchForm.Update(validationErr)
		standarizedResponse(c, true, http.StatusNotAcceptable, message, nil)
		return
	}

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleUpdate("sc_user.branch", form, cond, "id")
	if err != nil {
		fmt.Println(err)

		if pqErr, ok := err.(*pq.Error); ok {
			if strings.Contains(pqErr.Message, "duplicate key value violates unique constraint") {
				standarizedResponse(c, true, http.StatusConflict, "A branch with this name already exists", nil)
				return
			}
		}

		standarizedResponse(c, true, http.StatusNotAcceptable, "Branch could not be updated", nil)
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
		UserID: parsedUsrId, Type: "Branch", Detail: "Updated Branch",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Branch updated", nil)
}

// Delete ...
func (ctrl BranchController) Delete(c *gin.Context) {
	id := c.Param("id")

	cond := &interType.UpdateCond{
		Ids:  "id",
		Vals: id,
	}

	err := models.FlexibleDelete("sc_user.branch", cond)
	if err != nil {
		standarizedResponse(c, true, http.StatusNotAcceptable, "Branch could not be deleted", nil)
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
		UserID: parsedUsrId, Type: "Branch", Detail: "Deleted Branch",
	}
	go models.LogActivity(actLog)
	// ------------------------------

	standarizedResponse(c, false, http.StatusOK, "Branch deleted", nil)
}
