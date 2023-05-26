package controllers

import (
	"chatbox-api/pkg/secrets"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func ConvertStringToInt(s string) int {
	v, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return v
}

func HashPassword(password string) (string, error) {
	passwordHashCost := ConvertStringToInt(secrets.GetEnv("PASSWORD_HASH_COST"))

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)

	return string(bytes), err
}
