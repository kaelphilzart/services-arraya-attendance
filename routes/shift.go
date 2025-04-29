package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func shiftRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.ShiftController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	shift := rg.Group("/shift")

	// GET
	shift.GET("", tokenMid.ValidateAdmin(), controllers.All)
	shift.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	shift.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	shift.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	shift.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
