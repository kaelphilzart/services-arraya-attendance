package middlewares

import (
	"services-arraya-attendance/controllers"

	"github.com/gin-gonic/gin"
)

var authController = new(controllers.AuthController)

type TokenAuthMiddleware struct{}

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func (mid TokenAuthMiddleware) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authController.TokenValid(c)
		c.Next()
	}
}

func (mid TokenAuthMiddleware) ValidateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authController.TokenAdminValid(c)
		c.Next()
	}
}

// JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func (mid TokenAuthMiddleware) ValidateResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		authController.TokenResetPasswordValid(c)
		c.Next()
	}
}
