package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func positionRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.PositionController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	positions := rg.Group("/position")

	// GET
	positions.GET("", tokenMid.ValidateAdmin(), controllers.All)
	positions.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	positions.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	positions.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	positions.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
