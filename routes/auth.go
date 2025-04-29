package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func authRoutes(rg *gin.RouterGroup) {
	controllersUser := new(controllers.UserController)
	controllersAuth := new(controllers.AuthController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	users := rg.Group("/auth")

	// GET
	users.GET("/logout", tokenMid.Validate(), controllersUser.Logout)

	// POST
	users.POST("/login", controllersUser.Login)
	users.POST("/refresh", controllersAuth.Refresh)
}
