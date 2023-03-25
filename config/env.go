package config

import "fmt"

// "os"

type EnvVariables struct {
	DBCon string
}

var Env *EnvVariables

func InitializeEnv() {

	Env = &EnvVariables{
		DBCon: "hariprasanth:12345@tcp(127.0.0.1:3306)/sample",
		//DBCon: "root:12345@tcp(docker.for.mac.localhost:3306)/sample",
	}

	fmt.Println("Environment Variables Setup Done!!!")

}
