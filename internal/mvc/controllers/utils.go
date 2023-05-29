package controllers

import (
	"chatbox-api/pkg/secrets"
	"log"
	"reflect"
	"strconv"
	"strings"

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

func TrimWhiteSpacesStruct[T any](body *T) {
	values := reflect.ValueOf(body)
	bodyLength := reflect.Indirect(values).NumField()

	for i := 0; i < bodyLength; i++ {
		if reflect.Indirect(values).Field(i).Kind() == reflect.String {
			v := strings.TrimSpace(reflect.Indirect(values).Field(i).Interface().(string))
			reflect.ValueOf(body).Elem().Field(i).SetString(v)
		}
	}

}
