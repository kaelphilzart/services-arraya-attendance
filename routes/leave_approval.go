package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func leaveApprovalRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.LeaveApprovalController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	leaveApproval := rg.Group("/leaveApproval")

	// GET
	leaveApproval.GET("", tokenMid.ValidateAdmin(), controllers.All)
	leaveApproval.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// User
	leaveApproval.POST("", tokenMid.Validate(), controllers.Approval)

}
