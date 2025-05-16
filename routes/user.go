package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.UserController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	users := rg.Group("/user")

	// GET
	users.GET("", tokenMid.ValidateAdmin(), controllers.All)
	users.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	users.POST("", tokenMid.ValidateAdmin(), controllers.Create)
	users.POST("/createUser", controllers.CreateUser)

	// DELETE
	users.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)

	// PUT
	users.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)
}
