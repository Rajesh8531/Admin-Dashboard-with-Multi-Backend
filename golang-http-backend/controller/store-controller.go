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

var CreateStoreController = http.HandlerFunc(storeController)

func storeController(w http.ResponseWriter, r *http.Request) {

	var store types.Store
	var err error

	_ = utils.ParseJSON(r, &store)
	store.CreatedAt = time.Now()
	store.UpdatedAt = time.Now()
	var userId = r.Header.Get("id")

	store.UserId, _ = primitive.ObjectIDFromHex(userId)
	store.ID = primitive.NewObjectID()

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	_, err = db.Stores.InsertOne(context.Background(), store)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a store"))
		return
	}
	utils.ResponseJSON(w, http.StatusOK, store)
}

func GetStores(w http.ResponseWriter, r *http.Request) {

	var stores []types.Store

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	queries := r.URL.Query()

	userId, err := primitive.ObjectIDFromHex(queries["userId"][0])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	stores, err = db.GetStoresByValue(bson.M{"user_id": userId})
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, stores)
}

func GetStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseJSON(w, http.StatusAccepted, nil)
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)
	var store types.FullStoreType

	var st types.Store
	db.GetStoreById(&st, bson.M{"_id": storeId})
	store.ID = st.ID
	store.UpdatedAt = st.UpdatedAt
	store.CreatedAt = st.CreatedAt
	store.UserId = st.UserId
	store.Name = st.Name
	store.Billboards, _ = db.GetBillboardsByValue(bson.M{"storeId": storeId})
	store.Colors, _ = db.GetColorsByValue(bson.M{"storeId": storeId})
	store.Sizes, _ = db.GetSizesByValue(bson.M{"storeId": storeId})

	store.Categories, _ = db.GetCategoriesByValueData(bson.M{"storeId": storeId})
	utils.ResponseJSON(w, http.StatusAccepted, store)
}
