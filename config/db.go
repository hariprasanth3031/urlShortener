package config

import (
	"fmt"
	//"time"

	sql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func InitializeDb() {

	db, err := gorm.Open(sql.Open(Env.DBCon+fmt.Sprintf("?%s", "&parseTime=True")), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true, TablePrefix: "url_shortener."},
	})

	if err != nil {
		panic(err)
	}

	Db = db
	fmt.Println("Database Initialized!!!")
}
