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

var ProductsController = http.HandlerFunc(productsController)

func productsController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProducts(w, r)
		return
	} else {
		createProduct(w, r)
		return
	}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var productWithImage types.ProductWithImageUrl
	var product types.Product
	var err error

	_ = utils.ParseJSON(r, &productWithImage)

	params := mux.Vars(r)

	storeId, _ := primitive.ObjectIDFromHex(params["storeId"])

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	product.StoreId = storeId
	product.CategoryId = productWithImage.CategoryId
	product.SizeId = productWithImage.SizeId
	product.ColorId = productWithImage.ColorId
	product.Name = productWithImage.Name
	product.Price = productWithImage.Price
	product.IsArchived = productWithImage.IsArchived
	product.IsFeatured = productWithImage.IsFeatured
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.ID = primitive.NewObjectID()

	_, err = db.Products.InsertOne(context.Background(), product)
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	for _, url := range productWithImage.ImageUrl {
		var image types.Image

		image.CreatedAt = time.Now()
		image.UpdatedAt = time.Now()
		image.ProductId = product.ID
		image.Url = url

		_, err = db.Images.InsertOne(context.Background(), image)
	}

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a billboard"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, product)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var products []types.FullProduct
	products, err = db.GetProductsByValue(bson.M{"storeId": storeId})
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a products"))
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, products)
}

var UniqueProductController = http.HandlerFunc(uniqueProductController)

func uniqueProductController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		getProduct(w, r)
		return

	case "PATCH":
		updateProduct(w, r)
		return

	case "DELETE":
		deleteProduct(w, r)
		return
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	storeId, err := primitive.ObjectIDFromHex(params["storeId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	productId, err := primitive.ObjectIDFromHex(params["productId"])

	if err != nil {
		utils.ResponseJSON(w, http.StatusAccepted, nil)
		return
	}

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	var product types.FullProduct
	filter := bson.M{"_id": productId, "storeId": storeId}
	db.GetProduct(&product, filter)

	utils.ResponseJSON(w, http.StatusAccepted, product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var productWithImage types.ProductWithImageUrl
	var product types.Product
	var err error

	_ = utils.ParseJSON(r, &productWithImage)

	params := mux.Vars(r)

	productId, err := primitive.ObjectIDFromHex(params["productId"])
	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid cateogry_id"))
		return
	}

	storeId, _ := primitive.ObjectIDFromHex(params["storeId"])

	connection := db.ConnectToDB()
	defer db.CloseDB(connection)

	isAuthrorized := utils.IsAuthorizedForStore(r, storeId)

	if !isAuthrorized {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf(" unauthorized"))
		return
	}

	product.StoreId = storeId
	product.CategoryId = productWithImage.CategoryId
	product.SizeId = productWithImage.SizeId
	product.ColorId = productWithImage.ColorId
	product.Name = productWithImage.Name
	product.Price = productWithImage.Price
	product.IsArchived = productWithImage.IsArchived
	product.IsFeatured = productWithImage.IsFeatured
	product.CreatedAt = productWithImage.CreatedAt
	product.UpdatedAt = time.Now()
	product.ID = productId

	filter := bson.M{"_id": product.ID, "storeId": storeId}

	update := bson.M{"$set": product}

	_, err = db.Images.DeleteMany(context.Background(), bson.M{"productId": productId})

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	_, err = db.Products.UpdateOne(context.Background(), filter, update)
	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("invalid storeId"))
		return
	}

	for _, url := range productWithImage.ImageUrl {
		var image types.Image

		image.CreatedAt = time.Now()
		image.UpdatedAt = time.Now()
		image.ProductId = product.ID
		image.Url = url

		_, err = db.Images.InsertOne(context.Background(), image)
	}

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while inserting a billboard"))
		return
	}

	utils.ResponseJSON(w, http.StatusOK, product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
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

	productId, err := primitive.ObjectIDFromHex(params["productId"])

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	filter := bson.M{"_id": productId, "storeId": storeId}
	_, err = db.Products.DeleteOne(context.Background(), filter)

	if err != nil {
		utils.ResponseError(w, http.StatusNotAcceptable, fmt.Errorf("invalid store_id"))
		return
	}

	_, err = db.Images.DeleteMany(context.Background(), bson.M{"productId": productId})

	if err != nil {
		utils.ResponseError(w, http.StatusConflict, fmt.Errorf("error while fecthing a products"))
		return
	}
	utils.ResponseJSON(w, http.StatusAccepted, "DELETED SUCCESSFULLY")
}
