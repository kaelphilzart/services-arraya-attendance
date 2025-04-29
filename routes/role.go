package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func roleRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.RoleController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	role := rg.Group("/role")

	// GET
	role.GET("", tokenMid.ValidateAdmin(), controllers.All)
	role.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	role.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	role.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	role.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
