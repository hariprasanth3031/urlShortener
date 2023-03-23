package routes

import (
	"SD/handler"

	"github.com/gin-gonic/gin"
)

func createCheckRoutes(server *gin.RouterGroup) {
	check := server.Group("/short")
	{
		check.POST("/encode/url", handler.Encode)
		check.GET("/:url", handler.Decode)
	}
}
