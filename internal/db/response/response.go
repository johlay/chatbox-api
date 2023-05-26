package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

type SuccessResponse struct {
	Status  int
	Message string
	Data    interface{}
}

type UserResponse struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func HandleErrorResponse(w http.ResponseWriter, statusCode int, r ErrorResponse) {
	res := ErrorResponse{
		Status:  r.Status,
		Message: r.Message,
		Data:    r.Data,
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	return
}

func HandleSuccessResponse(w http.ResponseWriter, statusCode int, s SuccessResponse) {
	res := SuccessResponse{
		Status:  s.Status,
		Message: s.Message,
		Data:    s.Data,
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	return
}