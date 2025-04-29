package routes

import (
	"services-arraya-attendance/controllers"
	"services-arraya-attendance/middlewares"

	"github.com/gin-gonic/gin"
)

func branchRoutes(rg *gin.RouterGroup) {
	controllers := new(controllers.BranchController)
	tokenMid := new(middlewares.TokenAuthMiddleware)
	branches := rg.Group("/branch")

	// GET
	branches.GET("", tokenMid.ValidateAdmin(), controllers.All)
	branches.GET("/:id", tokenMid.ValidateAdmin(), controllers.One)

	// POST
	branches.POST("", tokenMid.ValidateAdmin(), controllers.Create)

	// PUT
	branches.PUT("/:id", tokenMid.ValidateAdmin(), controllers.Update)

	// DELETE
	branches.DELETE("/:id", tokenMid.ValidateAdmin(), controllers.Delete)
}
