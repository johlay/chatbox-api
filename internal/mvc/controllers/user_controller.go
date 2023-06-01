package controllers

import (
	"chatbox-api/internal/db"
	"chatbox-api/internal/db/model"
	"chatbox-api/internal/response"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.GetCollection(db.DB, "users")

/*
POST - login an user

returns HTTP response containing information on user's authentication status.
*/
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body response.LoginCredentialsResponse
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.HandleErrorResponse(w, 400, response.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Error occurred while decoding HTTP body response",
			Data:    map[string]interface{}{},
		})
		return
	} else {
		TrimWhiteSpacesStruct[response.LoginCredentialsResponse](&body)
	}

	filter := bson.D{{Key: "email", Value: body.Email}}

	if err := userCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		response.HandleErrorResponse(w, 403, response.ErrorResponse{
			Status:  http.StatusForbidden,
			Message: "Authentication was unsuccessful",
			Data:    map[string]interface{}{},
		})
		return
	}

	pwIsMatch := CheckPasswordHash(body.Password, user.Password)

	if pwIsMatch == false {
		response.HandleErrorResponse(w, 403, response.ErrorResponse{
			Status:  http.StatusForbidden,
			Message: "Authentication was unsuccessful",
			Data:    map[string]interface{}{},
		})
		return
	}

	payload := response.LoginSuccessResponse{
		ID:         user.ID,
		Email:      user.Email,
		First_name: user.First_name,
		Last_name:  user.Last_name,
	}

	jwtAccessToken, err := GenerateJWTAccessToken[response.LoginSuccessResponse](payload)

	if err != nil {
		log.Fatal(err)
	}

	response.HandleSuccessResponse(w, 200, response.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Authentication was successful",
		Data: struct {
			Access_token string                        `json:"access_token"`
			User         response.LoginSuccessResponse `json:"user"`
		}{
			Access_token: jwtAccessToken,
			User:         payload,
		},
	})
}

/*
POST - register a new user

returns HTTP response containing information on the newly registered user.
*/
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
	} else {
		TrimWhiteSpacesStruct[response.UserResponse](&body)
	}

	filter := bson.D{{Key: "email", Value: body.Email}}

	if err := userCollection.FindOne(context.TODO(), filter).Decode(&user); err == nil && body.Email == user.Email {
		response.HandleErrorResponse(w, 409, response.ErrorResponse{
			Status:  http.StatusConflict,
			Message: "The email you entered is already registered",
			Data:    map[string]interface{}{},
		})
		return
	}

	hashedPassword, err := HashPassword(body.Password)

	if err != nil {
		response.HandleErrorResponse(w, 500, response.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error occurred while hashing the pass",
			Data:    map[string]interface{}{},
		})
	}

	newUser := model.User{
		ID:         primitive.NewObjectID(),
		First_name: body.First_name,
		Last_name:  body.Last_name,
		Email:      body.Email,
		Password:   hashedPassword,
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
