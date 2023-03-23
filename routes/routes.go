package routes

import "github.com/gin-gonic/gin"

func CreateRoutes(server *gin.RouterGroup) {

	createCheckRoutes(server)
}
