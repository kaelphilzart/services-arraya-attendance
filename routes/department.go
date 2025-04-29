package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func departmentRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.DepartmentController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	departments := rg.Group("/department")

	// GET
	departments.GET("", tokenMid.ValidateAdmin(), controllers.All)
	departments.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	departments.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	departments.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	departments.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
