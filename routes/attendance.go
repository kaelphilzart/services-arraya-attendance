package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func attendanceRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.AttendanceController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	attendance := rg.Group("/attendance")

	//admin
	attendance.GET("", tokenMid.ValidateAdmin(), controllers.All)

	//user
	attendance.GET("/:id", tokenMid.Validate(), controllers.One)
	attendance.POST("/in", tokenMid.Validate(), controllers.AttendanceIn)
	attendance.PUT("/out", tokenMid.Validate(), controllers.AttendanceOut)
	attendance.GET("/history", tokenMid.Validate(), controllers.History)
	attendance.GET("/user", tokenMid.Validate(), controllers.OneByUserId)


}