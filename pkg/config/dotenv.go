package config

import (
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
)

func LoadDotEnv() {
	err := godotenvvault.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return
}

func GetEnv(v string) string {
	LoadDotEnv()

	val, ok := os.LookupEnv(v)

	if !ok {
		log.Fatal("Missing env variable: ", v)
	}

	return val
}
