package controllers

import (
	"fmt"
	"reflect"
	interType "services-arraya-attendance/interfaces"

	"github.com/gin-gonic/gin"
)

// Controller ...
type UtilsController struct{}

// StandarizedResponse ...
func standarizedResponse(c *gin.Context, isAbort bool, code int, message string, data interface{}) {
	response := &interType.StandardResponse{
		Status:  code,
		Message: message,
	}
	if data != nil {
		response.Data = data
	}
	if isAbort {
		c.AbortWithStatusJSON(code, response)
	} else {
		c.JSON(code, response)
	}
}

func mergeStructs(source interface{}, destination interface{}) {
	sourceValue := reflect.ValueOf(source)
	destinationValue := reflect.ValueOf(destination)

	if sourceValue.Kind() != reflect.Struct || destinationValue.Kind() != reflect.Ptr || destinationValue.Elem().Kind() != reflect.Struct {
		fmt.Println("Invalid input. Please provide structs.")
		return
	}

	for i := 0; i < sourceValue.NumField(); i++ {
		srcField := sourceValue.Field(i)
		dstField := destinationValue.Elem().FieldByName(sourceValue.Type().Field(i).Name)

		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}
}

// Includes checks if a slice contains a specific element
func CheckIncludes(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
