package routes

import (
	"urlshortener/handler"

	"github.com/gin-gonic/gin"
)

func createCheckRoutes(server *gin.RouterGroup) {

	//Grouping url shortening related roots under one roof
	check := server.Group("/short")
	{
		check.POST("/encode/url", handler.Encode)
		check.GET("/:url", handler.Decode)
	}
}
