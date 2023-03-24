package service

import (
	// "fmt"
	"net/http"
	"urlshortener/config"
	"urlshortener/database"
	"urlshortener/models"
)

func Decode(input string) models.JSONOutput {

	config.Logger.Debug("Service - Decode the tiny url")

	output, err := database.Decode(input)

	if err != nil {
		return models.JSONOutput{
			Code: http.StatusInternalServerError,
			Data: "Error while fetching the longurl",
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

	id, err := database.InsertLongUrl(input.Url)

	if err != nil {
		return models.JSONOutput{
			Code: http.StatusInternalServerError,
			Data: err,
		}
	}

	output, err = database.Encode(id)
	if err != nil {
		return models.JSONOutput{
			Code: http.StatusInternalServerError,
			Data: err,
		}
	}

	return models.JSONOutput{
		Code: http.StatusOK,
		Data: output,
	}

}
