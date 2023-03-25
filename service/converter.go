package service

import (
	"net/http"
	"urlshortener/config"
	"urlshortener/database"
	"urlshortener/models"
)

func Encode(input models.ConverterInput) models.JSONOutput {

	config.Logger.Debug("Service - Convert into short url")

	var output string
	var err error

	//Insert the long url into db and get the unique id
	id, err := database.InsertLongUrl(input.Url)

	if err != nil {
		return models.JSONOutput{
			Code: http.StatusInternalServerError,
			Data: err,
		}
	}

	//Encode the unique id and insert the short url into that id
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

func Decode(input string) models.JSONOutput {

	config.Logger.Debug("Service - Decode the tiny url")

	//Fetch the long url for the given short url
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
