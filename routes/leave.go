package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func leaveRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.LeaveController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	leave := rg.Group("/leave")

	// GET
	leave.GET("", tokenMid.ValidateAdmin(), controllers.All)
	leave.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// User
	leave.POST("", tokenMid.Validate(), controllers.Pengajuan)
	leave.GET("/department/:id", tokenMid.Validate(), controllers.OneByDepartment)

}
