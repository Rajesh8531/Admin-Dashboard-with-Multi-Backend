package controller

import (
	"connection-to-mongo/project/db"
	"connection-to-mongo/project/types"
	"connection-to-mongo/project/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ColorsController = http.HandlerFunc(colorsController)

func colorsController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getColors(w, r)
		return
	} else {
		createColor(w, r)
		return
	}
}

func createColor(w http.ResponseWriter, r *http.Request) {
	var color types.Color
	var err error

	_ = utils.ParseJSON(r, &color)

	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	color.StoreId = storeId
	color.CreatedAt = time.Now()
	color.UpdatedAt = time.Now()
	color.ID = primitive.NewObjectID()

	_, err = db.Colors.InsertOne(context.Background(), color)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a color"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, color)
}

func getColors(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var colors []types.Color
	colors, err = db.GetColorsByValue(bson.M{"storeId": storeId})
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a colors"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, colors)
}

var UniqueColorController = http.HandlerFunc(uniqueColorController)

func uniqueColorController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		getColor(w, r)
		return

	case "PATCH":
		updateColor(w, r)
		return

	case "DELETE":
		deleteColor(w, r)
		return
	}
}

func getColor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	colorId, err := primitive.ObjectIDFromHex(params["colorId"])

	if err != nil {
		utils.ResponseJSON(w, http.StatusAccepted, nil)
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var color types.Color
	filter := bson.M{"_id": colorId, "storeId": storeId}
	db.GetColorById(&color, filter)

	utils.ResponseJSON(w, http.StatusAccepted, color)
}

func updateColor(w http.ResponseWriter, r *http.Request) {
	var newColor types.Color
	params := mux.Vars(r)
	utils.ParseJSON(r, &newColor)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	newColor.StoreId = storeId

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	colorId, err := primitive.ObjectIDFromHex(params["colorId"])
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid cateogry_id"))
		return
	}
	newColor.ID = colorId
	newColor.UpdatedAt = time.Now()

	update := bson.M{"$set": newColor}

	filter := bson.M{"_id": colorId, "storeId": newColor.StoreId}
	_, err = db.Colors.UpdateOne(context.Background(), filter, update)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, newColor)
}

func deleteColor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	colorId, err := primitive.ObjectIDFromHex(params["colorId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	filter := bson.M{"_id": colorId, "storeId": storeId}
	_, err = db.Colors.DeleteOne(context.Background(), filter)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}
	utils.ResponseJSON(w, http.StatusAccepted, "DELETED SUCCESSFULLY")
}
