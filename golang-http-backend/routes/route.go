package routes

import (
	"connection-to-mongo/project/controller"
	"connection-to-mongo/project/middleware"

	"github.com/gorilla/mux"
)

func AddAuthorizationRoute(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", controller.SignUp).Methods("POST")
	auth.HandleFunc("/signin", controller.SignIn).Methods("POST")

}

func AddStoreRoute(r *mux.Router) {
	store := r.PathPrefix("/store").Subrouter()
	store.Handle("/", middleware.AuthMiddleware(controller.CreateStoreController)).Methods("POST")
	store.HandleFunc("/", controller.GetStores).Methods("GET")
	store.HandleFunc("/{storeId}", controller.GetStore).Methods("GET")

}

func AddBillboardRoute(r *mux.Router) {

	billboard := r.PathPrefix("/store/{storeId}/billboards").Subrouter()
	billboard.Handle("", middleware.AuthMiddleware(controller.BillboardsController)).Methods("POST", "GET")
	billboard.Handle("/{billboardId}", middleware.AuthMiddleware(controller.UniqueBillboardController)).Methods("GET", "PATCH", "DELETE")

}

func AddCategoryRoute(r *mux.Router) {

	billboard := r.PathPrefix("/store/{storeId}/categories").Subrouter()
	billboard.Handle("", middleware.AuthMiddleware(controller.CategoriesController)).Methods("POST", "GET")
	billboard.Handle("/{categoryId}", middleware.AuthMiddleware(controller.UniqueCategoryController)).Methods("GET", "PATCH", "DELETE")

}

func AddSizeRoute(r *mux.Router) {

	billboard := r.PathPrefix("/store/{storeId}/sizes").Subrouter()
	billboard.Handle("", middleware.AuthMiddleware(controller.SizesController)).Methods("POST", "GET")
	billboard.Handle("/{sizeId}", middleware.AuthMiddleware(controller.UniqueSizeController)).Methods("GET", "PATCH", "DELETE")

}

func AddColorRoute(r *mux.Router) {

	billboard := r.PathPrefix("/store/{storeId}/colors").Subrouter()
	billboard.Handle("", middleware.AuthMiddleware(controller.ColorsController)).Methods("POST", "GET")
	billboard.Handle("/{colorId}", middleware.AuthMiddleware(controller.UniqueColorController)).Methods("GET", "PATCH", "DELETE")

}

func AddProductRoute(r *mux.Router) {

	billboard := r.PathPrefix("/store/{storeId}/products").Subrouter()
	billboard.Handle("", middleware.AuthMiddleware(controller.ProductsController)).Methods("POST", "GET")
	billboard.Handle("/{productId}", middleware.AuthMiddleware(controller.UniqueProductController)).Methods("GET", "PATCH", "DELETE")

}

func AddOrderRoute(r *mux.Router) {

	billboard := r.PathPrefix("/store/{storeId}/orders").Subrouter()
	billboard.Handle("", middleware.AuthMiddleware(controller.OrdersController)).Methods("POST", "GET")
	// billboard.Handle("/{orderId}", middleware.AuthMiddleware(controller.UniqueProductController)).Methods("GET", "PATCH", "DELETE")
}
