package main

import (
	"fmt"
	"net/http"
	"urlshortener/config"
	"urlshortener/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Url Shortening service")

	//Intitialize configurations
	config.InitializeLogger()
	config.InitializeEnv()
	config.InitializeDb()

	//Starting Gin server
	server := gin.New()

	group := server.Group("/tinyurl")

	routes.CreateRoutes(group)

	//Health check url
	group.GET("/public/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	//listens at 8080
	fmt.Println("Running at port: 8080")
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
