package controller

import (
	"connection-to-mongo/project/db"
	"connection-to-mongo/project/types"
	"connection-to-mongo/project/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var user types.User

	err := utils.ParseJSON(r, &user)

	if err != nil {
		utils.ResponseError(w, http.StatusExpectationFailed, err)
		return
	}

	var existingUser types.User
	_ = db.Users.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&existingUser)
	if existingUser.Email != "" {
		utils.ResponseError(w, http.StatusAlreadyReported, fmt.Errorf("user already existing"))
		return
	}

	user.Password, err = utils.GeneratehashPassword(user.Password)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error in password generation"))
		return
	}

	_, err = db.Users.InsertOne(context.Background(), user)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error in creating a new user"))
		return
	}

	tokenString, err := utils.GenerateJWT(user.Email, user.ID.String())

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error in jwt token generation"))
		return
	}

	response := types.TokenResponse{
		Token: tokenString,
		Name:  user.Name,
		Email: user.Email,
		ID:    user.ID,
	}

	utils.ResponseJSON(w, http.StatusOK, response)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var user types.User

	err := utils.ParseJSON(r, &user)

	if err != nil {
		utils.ResponseError(w, http.StatusExpectationFailed, err)
		return
	}

	var existingUser types.User
	_ = db.Users.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&existingUser)
	if existingUser.Email == "" {
		utils.ResponseError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	isPasswordCorrect := utils.CheckPasswordHash(user.Password, existingUser.Password)

	if !isPasswordCorrect {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid credentials"))
		return
	}

	tokenString, err := utils.GenerateJWT(existingUser.Email, existingUser.ID.String())

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error in jwt token generation"))
		return
	}

	response := types.TokenResponse{
		Token: tokenString,
		Name:  existingUser.Name,
		Email: existingUser.Email,
		ID:    existingUser.ID,
	}
	utils.ResponseJSON(w, http.StatusOK, response)
}
