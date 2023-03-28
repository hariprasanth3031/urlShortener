package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type EnvVariables struct {
	//DBCon string
	USER   string
	PASS   string
	HOST   string
	DBNAME string
}

var Env *EnvVariables

func InitializeEnv() {

	_ = godotenv.Load(".env")
	Env = &EnvVariables{
		USER:   os.Getenv("USER"),
		PASS:   os.Getenv("PASSWORD"),
		HOST:   os.Getenv("HOST"),
		DBNAME: os.Getenv("NAME"),
	}

	v := validator.New()
	err := v.Struct(*Env)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}
