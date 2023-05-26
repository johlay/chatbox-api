package mvc

import (
	"chatbox-api/internal/db"
	"chatbox-api/internal/db/model"
	"chatbox-api/internal/response"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.GetCollection(db.DB, "users")

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var body response.UserResponse
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.HandleErrorResponse(w, 400, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "error while parsing http request body object: " + err.Error(),
			Data:    map[string]interface{}{},
		})
		return
	}

	filter := bson.D{{Key: "email", Value: body.Email}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err == nil && body.Email == user.Email {
		response.HandleErrorResponse(w, 409, response.ErrorResponse{
			Status:  http.StatusConflict,
			Message: "The email you entered is already registered",
			Data:    map[string]interface{}{},
		})
		return
	}

	newUser := model.User{
		ID:         primitive.NewObjectID(),
		First_name: body.First_name,
		Last_name:  body.Last_name,
		Email:      body.Email,
		Password:   body.Password,
	}

	result, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		response.HandleErrorResponse(w, 500, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    map[string]interface{}{},
		})
		return
	}

	response.HandleSuccessResponse(w,
		http.StatusCreated,
		response.SuccessResponse{
			Status:  201,
			Message: "Successfully created a new user",
			Data:    result,
		},
	)
}
