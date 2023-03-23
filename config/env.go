package config

import (
// "os"
)

type EnvVariables struct {
	DBCon string
}

var Env *EnvVariables

func InitializeEnv() {

	Env = &EnvVariables{
		DBCon: "hariprasanth:12345@tcp(127.0.0.1:3306)/sample",
	}

}
