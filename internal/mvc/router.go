package mvc

import (
	"chatbox-api/internal/mvc/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()
var PORT = os.Getenv("PORT")

func userRoutes() {
	s := Router.PathPrefix("/api/user").Subrouter()

	s.HandleFunc("/login", controllers.HandleLogin)
	s.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
}

func SetupRouter() {
	userRoutes()

	http.ListenAndServe(":"+PORT, Router)
}
