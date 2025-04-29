package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("ping")

	// GET
	ping.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong!")
	})
}
