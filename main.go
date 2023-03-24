package main

import (
	"fmt"
	"net/http"

	//"/Users/hari/Desktop/SD/routes"
	"SD/config"
	"SD/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")

	config.InitializeLogger()
	config.InitializeEnv()
	config.InitializeDb()

	server := gin.New()

	group := server.Group("/tinyurl")

	routes.CreateRoutes(group)

	group.GET("/public/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	//listens at 8003
	if err := server.Run(":8003"); err != nil {
		panic(err)
	}
}
