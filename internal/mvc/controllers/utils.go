package controllers

import (
	"chatbox-api/pkg/config"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var JWT_ACCESS_TOKEN_SECRET = config.GetEnv("JWT_ACCESS_TOKEN_SECRET")
var JWT_ACCESS_TOKEN_EXPIRY_TIME_HOUR = ConvertStringToInt(config.GetEnv("JWT_ACCESS_TOKEN_EXPIRY_TIME_HOUR"))

type JWTCustomClaims[T any] struct {
	Payload T `json:"access_token"`
	jwt.RegisteredClaims
}

func CheckPasswordHash(password, dbPasswordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password))

	return err == nil
}

func ConvertStringToInt(s string) int {
	v, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return v
}

func GenerateJWTAccessToken[T any](payload T) (string, error) {
	claims := JWTCustomClaims[T]{payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(JWT_ACCESS_TOKEN_EXPIRY_TIME_HOUR))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(JWT_ACCESS_TOKEN_SECRET))

	return ss, err
}

func HashPassword(password string) (string, error) {
	passwordHashCost := ConvertStringToInt(config.GetEnv("PASSWORD_HASH_COST"))

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
