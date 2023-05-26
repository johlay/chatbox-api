package mvc

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/user").Subrouter()

	s.HandleFunc("/login", HandleLogin)
	s.HandleFunc("/register", HandleRegister).Methods("POST")

	http.ListenAndServe(":8080", r)
}
