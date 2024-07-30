package db

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Users *mongo.Collection
var Stores *mongo.Collection
var Billboards *mongo.Collection
var Categories *mongo.Collection
var Colors *mongo.Collection
var Images *mongo.Collection
var Orders *mongo.Collection
var OrderItems *mongo.Collection
var Products *mongo.Collection
var Sizes *mongo.Collection

func ConnectToDB() *mongo.Client {

	envFile, _ := godotenv.Read()
	clientOptions := options.Client().ApplyURI(envFile["MONGO_CONNECTION_URI"])
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SUCCESSFULLY CONNECTED TO DATABASE")

	Users = client.Database("admin-dashboard").Collection("users")
	Stores = client.Database("admin-dashboard").Collection("Store")
	Billboards = client.Database("admin-dashboard").Collection("Billboard")
	Categories = client.Database("admin-dashboard").Collection("Category")
	Colors = client.Database("admin-dashboard").Collection("Color")
	Images = client.Database("admin-dashboard").Collection("Image")
	Orders = client.Database("admin-dashboard").Collection("Order")
	OrderItems = client.Database("admin-dashboard").Collection("OrderItem")
	Products = client.Database("admin-dashboard").Collection("Product")
	Sizes = client.Database("admin-dashboard").Collection("Size")

	return client
}

func CloseDB(client *mongo.Client) {
	client.Disconnect(context.TODO())
}
