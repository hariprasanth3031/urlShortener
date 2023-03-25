package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/config"
	"urlshortener/routes"
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

	//listens at 8003
	fmt.Println("Running at port: 8003")
	if err := server.Run(":8003"); err != nil {
		panic(err)
	}
}
