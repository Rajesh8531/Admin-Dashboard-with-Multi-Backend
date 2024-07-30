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

var BillboardsController = http.HandlerFunc(billboardsController)

func billboardsController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getBillboards(w, r)
		return
	}

	var billboard types.Billboard
	var err error

	_ = utils.ParseJSON(r, &billboard)

	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	billboard.StoreId = storeId
	billboard.CreatedAt = time.Now()
	billboard.UpdatedAt = time.Now()
	billboard.ID = primitive.NewObjectID()

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	_, err = db.Billboards.InsertOne(context.Background(), billboard)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a billboard"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, billboard)
}

func getBillboards(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var billboards []types.Billboard
	billboards, err = db.GetBillboardsByValue(bson.M{"storeId": storeId})
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, billboards)
}

func GetBillboard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	billboardId, err := primitive.ObjectIDFromHex(params["billboardId"])

	if err != nil {
		utils.ResponseJSON(w, http.StatusAccepted, nil)
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var billboard types.FullBillboardType
	var bb types.Billboard
	var categories []types.Category
	filter := bson.M{"_id": billboardId, "storeId": storeId}
	db.GetBillboardById(&bb, filter)

	billboard.ID = bb.ID
	billboard.CreatedAt = bb.CreatedAt
	billboard.UpdatedAt = bb.UpdatedAt
	billboard.Label = bb.Label
	billboard.StoreId = bb.StoreId
	billboard.ImageUrl = bb.ImageUrl
	billboard.Categories = categories

	utils.ResponseJSON(w, http.StatusAccepted, billboard)
}

var UniqueBillboardController = http.HandlerFunc(uniqueBillboardController)

func uniqueBillboardController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetBillboard(w, r)
		return
	case "PATCH":
		updateBillboard(w, r)
		return
	case "DELETE":
		deleteBillboard(w, r)
		return
	}
}

func updateBillboard(w http.ResponseWriter, r *http.Request) {
	var newBillboard types.Billboard
	params := mux.Vars(r)
	utils.ParseJSON(r, &newBillboard)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	newBillboard.StoreId = storeId

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
	}

	billboardId, err := primitive.ObjectIDFromHex(params["billboardId"])
	newBillboard.ID = billboardId
	newBillboard.UpdatedAt = time.Now()

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	update := bson.M{"$set": newBillboard}

	filter := bson.M{"_id": billboardId, "storeId": storeId}
	_, err = db.Billboards.UpdateOne(context.Background(), filter, update)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, newBillboard)
}

func deleteBillboard(w http.ResponseWriter, r *http.Request) {
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

	billboardId, err := primitive.ObjectIDFromHex(params["billboardId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	var billboard types.Billboard
	filter := bson.M{"_id": billboardId, "storeId": storeId}
	_, err = db.Billboards.DeleteOne(context.Background(), filter)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}
	utils.ResponseJSON(w, http.StatusAccepted, billboard)
}
