package database

import (
	"fmt"
	"time"
	"urlshortener/config"
	"urlshortener/constants"
	"urlshortener/helper"
	"urlshortener/models"
)

func Decode(input string) (string, error) {

	config.Logger.Debug("Database - Decode the tiny url")

	var db = config.Db

	var output string

	ctime := time.Now()

	if err := db.Debug().Raw("select long_url from url_store where short_url = ? and expires_at >= ?", input, ctime).
		Find(&output).Error; err != nil {
		return output, err
	}

	return output, nil
}

func Encode(n uint64) (string, error) {

	config.Logger.Debug("Database - Convert into tiny url")

	id := n

	var shorturl string

	for n > 0 {
		shorturl = shorturl + string(constants.AlphaNumerals[n%62])
		n /= 62
	}

	shorturl = helper.Reverse(shorturl)

	db := config.Db

	var input models.UrlStore

	input.ShortUrl = shorturl

	if err := db.Debug().Where("id = ?", id).Updates(input).Error; err != nil {
		return "", err
	}

	shorturl = "localhost:8003/tinyurl/short/" + shorturl

	return shorturl, nil
}

func InsertLongUrl(longurl string) uint64 {

	db := config.Db

	var input models.UrlStore
	input.LongUrl = longurl

	var id uint64

	times := time.Now().AddDate(0, 0, 1)
	ctime := time.Now()

	if res := db.Debug().Raw("Insert into url_store (long_url, created_at, expires_at) VALUES (?, ?, ?) RETURNING id", longurl, ctime, times).Scan(&id).Error; res != nil {
		return 0
	}

	fmt.Println(id)

	return id

}
