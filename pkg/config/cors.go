package config

import (
	"net/http"

	"github.com/rs/cors"
)

func EnableCors() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Content-Length"},
	})

	return c
}
