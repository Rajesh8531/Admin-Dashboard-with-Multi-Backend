package main

import (
	"connection-to-mongo/project/db"
	"connection-to-mongo/project/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Client
}

func (a *App) Initialize() {

	r := mux.NewRouter()

	routes.AddAuthorizationRoute(r)
	routes.AddStoreRoute(r)
	routes.AddBillboardRoute(r)
	routes.AddCategoryRoute(r)
	routes.AddSizeRoute(r)
	routes.AddColorRoute(r)
	routes.AddProductRoute(r)
	routes.AddOrderRoute(r)

	a.Router = r
}

func (a *App) Run(addr string) {
	fmt.Println("SERVER RUNNING ON ", addr)
	http.ListenAndServe(addr, cors.AllowAll().Handler(a.Router))
}

func main() {
	app := App{}
	app.Initialize()
	app.Run("localhost:3000")
	connection := db.ConnectToDB()
	defer db.CloseDB(connection)
}
