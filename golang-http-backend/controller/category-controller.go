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

var CategoriesController = http.HandlerFunc(categoriesController)

func categoriesController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCategories(w, r)
		return
	} else {
		createCategory(w, r)
		return
	}
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var category types.Category
	var err error

	_ = utils.ParseJSON(r, &category)

	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	category.StoreId = storeId
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	category.ID = primitive.NewObjectID()

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	_, err = db.Categories.InsertOne(context.Background(), category)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a billboard"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, category)
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var categories []types.FullCategory
	categories, err = db.GetCategoriesByValue(bson.M{"storeId": storeId})
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a categories"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, categories)
}

var UniqueCategoryController = http.HandlerFunc(uniqueCategoryController)

func uniqueCategoryController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		getCategory(w, r)
		return

	case "PATCH":
		updateCategory(w, r)
		return

	case "DELETE":
		deleteCategory(w, r)
		return
	}
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	categoryId, err := primitive.ObjectIDFromHex(params["categoryId"])

	if err != nil {
		utils.ResponseJSON(w, http.StatusAccepted, nil)
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	filter := bson.M{"_id": categoryId, "storeId": storeId}
	category := db.GetCategory(filter)

	utils.ResponseJSON(w, http.StatusAccepted, category)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory types.Category
	params := mux.Vars(r)
	utils.ParseJSON(r, &newCategory)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])
	newCategory.StoreId = storeId

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	categoryId, err := primitive.ObjectIDFromHex(params["categoryId"])
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid cateogry_id"))
	}
	newCategory.ID = categoryId
	newCategory.UpdatedAt = time.Now()

	update := bson.M{"$set": newCategory}

	filter := bson.M{"_id": categoryId, "storeId": storeId}
	_, err = db.Categories.UpdateOne(context.Background(), filter, update)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, newCategory)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
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

	categoryId, err := primitive.ObjectIDFromHex(params["categoryId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	filter := bson.M{"_id": categoryId, "storeId": storeId}
	_, err = db.Categories.DeleteOne(context.Background(), filter)

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a billboards"))
		return
	}
	utils.ResponseJSON(w, http.StatusAccepted, "DELETED SUCCESSFULLY")
}
