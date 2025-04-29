package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func userProfileRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.UserProfileController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	usersProfile := rg.Group("/user-profile")

	// GET
	usersProfile.GET("/:id", tokenMid.Validate(), controllers.One)

	// Update
	usersProfile.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)
}