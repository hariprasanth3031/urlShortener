package database

import (
	"SD/config"
	"SD/constants"
	"SD/models"
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomUrl(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CheckUrl(input string) bool {

	config.Logger.Debug("Database - Check duplicates")

	var db = config.Db

	var out string
	res := db.Debug().Select("short_url").Table("url_store").
		Where("short_url = ?", input).
		Find(&out)

	if res.RowsAffected > 0 {
		return true
	}
	return false
}

func Decode(input string) (string, error) {

	config.Logger.Debug("Database - Decode the tiny url")

	var db = config.Db

	var output string

	if res := db.Debug().Raw("select long_url from url_store where short_url = ?", input).
		Find(&output).Error; res != nil {
		return output, res
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

	db := config.Db

	var input models.UrlStore

	input.ShortUrl = shorturl

	if err := db.Debug().Where("id = ?", id).Updates(input).Error; err != nil {
		return "", nil
	}

	return shorturl, nil
}

func InsertLongUrl(longurl string) uint64 {

	db := config.Db

	var input models.UrlStore
	input.LongUrl = longurl

	var id uint64

	if res := db.Debug().Raw("Insert into url_store (long_url) VALUES (?) RETURNING id", longurl).Scan(&id).Error; res != nil {
		return 0
	}

	fmt.Println(id)

	return id

}
