package config

import (
	"fmt"
	//"time"

	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func InitializeDb() {

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s", Env.USER, Env.PASS, Env.HOST, Env.DBNAME)

	fmt.Println("url is", URL)

	db, err := gorm.Open(mysql.Open(URL+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "url_shortener.",
		},
	})

	if err != nil {
		panic(err)
	}

	Db = db
	fmt.Println("Database Initialized!!!")
}
