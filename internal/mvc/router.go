package mvc

import (
	"chatbox-api/internal/mvc/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func userRoutes() {
	s := Router.PathPrefix("/api/user").Subrouter()

	s.HandleFunc("/login", controllers.HandleLogin)
	s.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
}

func SetupRouter() {
	userRoutes()

	http.ListenAndServe(":8080", Router)
}
