package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func typeLeaveRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.TypeLeaveController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	typeLeave := rg.Group("/type-leave")

	// GET
	typeLeave.GET("", tokenMid.ValidateAdmin(), controllers.All)
	typeLeave.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	typeLeave.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	typeLeave.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	typeLeave.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
