package handler

import (
	"net/http"
	"urlshortener/config"
	"urlshortener/models"
	"urlshortener/service"

	"github.com/gin-gonic/gin"
)

func Encode(c *gin.Context) {

	config.Logger.Debug("Handler - Convert into short url")

	var input models.ConverterInput

	err := c.BindJSON(&input)

	if err != nil {
		config.Logger.Error("Error while binding the input")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Input"})
		return
	}

	output := service.Encode(input)

	c.JSON(output.Code, output.Data)
}

func Decode(c *gin.Context) {

	config.Logger.Debug("Handler - Decode the tinyurl")

	tinyUrl := c.Param("url")

	output := service.Decode(tinyUrl)

	c.JSON(output.Code, output.Data)
}
