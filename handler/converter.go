package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/config"
	"urlshortener/models"
	"urlshortener/service"
)

// Function to convert the long url into a tiny url
func Encode(c *gin.Context) {

	config.Logger.Debug("Handler - Convert into short url")

	//Binding the input with the model
	var input models.ConverterInput

	err := c.BindJSON(&input)

	if err != nil || input.Url == "" {
		config.Logger.Error("Error while binding the input")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Input"})
		return
	}

	output := service.Encode(input)

	c.JSON(output.Code, output.Data)
}

// Function to convert a tiny url to respective long url
func Decode(c *gin.Context) {

	config.Logger.Debug("Handler - Decode the tinyurl")

	//Get the tiny url by parsing the url
	tinyUrl := c.Param("url")

	output := service.Decode(tinyUrl)

	c.JSON(output.Code, output.Data)
}
