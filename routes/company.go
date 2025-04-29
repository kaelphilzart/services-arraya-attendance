package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.CompanyController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	Company := rg.Group("/company")

	// GET
	Company.GET("", tokenMid.ValidateAdmin(), controllers.All)
	Company.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	Company.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	Company.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	Company.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
