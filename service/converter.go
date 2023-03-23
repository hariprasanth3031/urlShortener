package service

import (
	// "fmt"
	"SD/config"
	"SD/database"
	"SD/models"
	"net/http"
)

func Decode(input string) models.JSONOutput {

	config.Logger.Debug("Service - Decode the tiny url")

	output, err := database.Decode(input)

	if err != nil {
		return models.JSONOutput{
			Code: http.StatusInternalServerError,
			Data: output,
		}
	}

	return models.JSONOutput{
		Code: http.StatusOK,
		Data: output,
	}

}

func Encode(input models.ConverterInput) models.JSONOutput {

	config.Logger.Debug("Service - Convert into short url")

	var output string
	var err error

	//shortenedUrl := database.GenerateRandomUrl(5)

	//IsPresent := database.CheckUrl(shortenedUrl)

	// if !IsPresent {

	//Insert the long url and get id

	id := database.InsertLongUrl(input.Url)

	output, err = database.Encode(id)
	if err != nil {
		return models.JSONOutput{
			Code: http.StatusInternalServerError,
			Data: output,
		}
	}
	// }

	return models.JSONOutput{
		Code: http.StatusOK,
		Data: output,
	}

}
