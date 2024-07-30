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

var OrdersController = http.HandlerFunc(ordersController)

func ordersController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getOrders(w, r)
		return
	} else {
		createOrder(w, r)
		return
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order types.Order
	var err error

	_ = utils.ParseJSON(r, &order)

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

	order.StoreId = storeId
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.ID = primitive.NewObjectID()

	_, err = db.Orders.InsertOne(context.Background(), order)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a billboard"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var orders []types.Order
	orders, err = db.GetOrdersByValue(bson.M{"storeId": storeId})
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a orders"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, orders)
}
