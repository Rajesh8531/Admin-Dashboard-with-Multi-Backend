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

var SizesController = http.HandlerFunc(sizesController)

func sizesController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getSizes(w, r)
		return
	} else {
		createSize(w, r)
		return
	}
}

func createSize(w http.ResponseWriter, r *http.Request) {
	var size types.Size
	var err error

	_ = utils.ParseJSON(r, &size)

	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}
	size.StoreId = storeId
	size.CreatedAt = time.Now()
	size.UpdatedAt = time.Now()
	size.ID = primitive.NewObjectID()

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	_, err = db.Sizes.InsertOne(context.Background(), size)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a billboard"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, size)
}

func getSizes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var sizes []types.Size
	sizes, err = db.GetSizesByValue(bson.M{"storeId": storeId})
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a sizes"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, sizes)
}

var UniqueSizeController = http.HandlerFunc(uniqueSizeController)

func uniqueSizeController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		getSize(w, r)
		return

	case "PATCH":
		updateSize(w, r)
		return

	case "DELETE":
		deleteSize(w, r)
		return
	}
}

func getSize(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	sizeId, err := primitive.ObjectIDFromHex(params["sizeId"])

	if err != nil {
		utils.ResponseJSON(w, http.StatusAccepted, nil)
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var size types.Size
	filter := bson.M{"_id": sizeId, "storeId": storeId}
	db.GetSizeById(&size, filter)

	utils.ResponseJSON(w, http.StatusAccepted, size)
}

func updateSize(w http.ResponseWriter, r *http.Request) {
	var newSize types.Size
	params := mux.Vars(r)
	utils.ParseJSON(r, &newSize)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	newSize.StoreId = storeId

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

	sizeId, err := primitive.ObjectIDFromHex(params["sizeId"])
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid cateogry_id"))
		return
	}
	newSize.ID = sizeId
	newSize.UpdatedAt = time.Now()

	update := bson.M{"$set": newSize}

	filter := bson.M{"_id": newSize.ID, "storeId": newSize.StoreId}
	_, err = db.Sizes.UpdateOne(context.Background(), filter, update)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, newSize)
}

func deleteSize(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	sizeId, err := primitive.ObjectIDFromHex(params["sizeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	filter := bson.M{"_id": sizeId, "storeId": storeId}
	_, err = db.Sizes.DeleteOne(context.Background(), filter)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}
	utils.ResponseJSON(w, http.StatusAccepted, "DELETED SUCCESSFULLY")
}
