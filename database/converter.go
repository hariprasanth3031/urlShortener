package database

import (
	"errors"
	"time"
	"urlshortener/config"
	"urlshortener/constants"
	"urlshortener/helper"
	"urlshortener/models"
)

func Encode(uniqueId uint64) (string, error) {

	config.Logger.Debug("Database - Convert into tiny url")

	uniqueIdCopy := uniqueId

	var shorturl string

	//Hashing the id using custom algorithm
	for uniqueIdCopy > 0 {
		shorturl = shorturl + string(constants.AlphaNumerals[uniqueIdCopy%62])
		uniqueIdCopy /= 62
	}

	//Reverse the generated url
	shorturl = helper.Reverse(shorturl)

	db := config.Db

	var input models.UrlStore

	input.ShortUrl = shorturl

	//Insert the short url into the db
	if err := db.Debug().Where("id = ?", uniqueId).Updates(input).Error; err != nil {
		return "", err
	}

	//Return the short url appended with base url
	shorturl = constants.BaseUrl + "/tinyurl/short/" + shorturl

	return shorturl, nil
}

func Decode(input string) (string, error) {

	config.Logger.Debug("Database - Decode the tiny url")

	var db = config.Db

	var output string

	//Fetch the current time
	currentTime := time.Now()

	if err := db.Debug().Raw("select long_url from url_store where short_url = ? and expires_at >= ?", input, currentTime).
		Find(&output).Error; err != nil {
		return output, err
	}

	return output, nil
}

func InsertLongUrl(longurl string) (uint64, error) {

	config.Logger.Debug("Database - Insert the long url")

	db := config.Db

	var input models.UrlStore
	input.LongUrl = longurl

	var id uint64

	//Set the expiry time
	times := time.Now().AddDate(0, 0, 1)
	ctime := time.Now()

	//Insert the longurl into db
	if res := db.Debug().Raw("INSERT INTO url_store (long_url, created_at, expires_at) VALUES (?, ?, ?) RETURNING id", longurl, ctime, times).Scan(&id).Error; res != nil {
		return 0, errors.New("unable to insert the url")
	}

	return id, nil

}
