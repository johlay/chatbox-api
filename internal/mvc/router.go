package mvc

import (
	"chatbox-api/internal/mvc/controllers"
	"chatbox-api/pkg/config"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()
var PORT = os.Getenv("PORT")

func userRoutes() {
	s := Router.PathPrefix("/api/user").Subrouter()

	s.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	s.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
}

func SetupRouter() {
	userRoutes()

	cors := config.EnableCors()

	http.ListenAndServe(":"+PORT, cors.Handler(Router))
}
