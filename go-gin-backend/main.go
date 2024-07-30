package main

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/middleware"
	"admin-dashboard/backend/golan-gin/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	db.ConnectToDB()
	defer db.CloseDB(db.DB)

	db.InitTables()

	auth := router.Group("/auth")
	{
		auth.POST("/signin", service.SignIn)
		auth.POST("/signup", service.SignUp)
	}

	store := router.Group("/store")
	{
		store.GET("/", service.GetStores)
		store.POST("/", middleware.Authorized, service.CreateStore)
		store.PATCH("/:storeId", middleware.Authorized, service.UpdateStore)
		store.GET("/:storeId", service.GetStore)
		store.DELETE("/:storeId", middleware.Authorized, service.DeleteStore)

		uniqueStore := store.Group("/:storeId")
		{
			billboard := uniqueStore.Group("/billboards")
			{
				billboard.POST("", middleware.Authorized, service.CreateBillboard)
				billboard.GET("", service.GetBillboards)
				billboard.GET("/:billboardId", service.GetBillboard)
				billboard.PATCH("/:billboardId", middleware.Authorized, service.UpdateBillboard)
				billboard.DELETE("/:billboardId", middleware.Authorized, service.DeleteBillboard)

			}

			category := uniqueStore.Group("/categories")
			{
				category.POST("", middleware.Authorized, service.CreateCategory)
				category.GET("", service.GetCategories)
				category.GET("/:categoryId", service.GetCategory)
				category.PATCH("/:categoryId", middleware.Authorized, service.UpdateCategory)
				category.DELETE("/:categoryId", middleware.Authorized, service.DeleteCategory)

			}

			sizes := uniqueStore.Group("/sizes")
			{
				sizes.POST("", middleware.Authorized, service.CreateSize)
				sizes.GET("", service.GetSizes)
				sizes.GET("/:sizeId", service.GetSize)
				sizes.PATCH("/:sizeId", middleware.Authorized, service.UpdateSize)
				sizes.DELETE("/:sizeId", middleware.Authorized, service.DeleteSize)

			}

			colors := uniqueStore.Group("/colors")
			{
				colors.POST("", middleware.Authorized, service.CreateColor)
				colors.GET("", service.GetColors)
				colors.GET("/:colorId", service.GetColor)
				colors.PATCH("/:colorId", middleware.Authorized, service.UpdateColor)
				colors.DELETE("/:colorId", middleware.Authorized, service.DeleteColor)

			}

			products := uniqueStore.Group("/products")
			{
				products.POST("", middleware.Authorized, service.CreateProduct)
				products.GET("", service.GetProducts)
				products.GET("/:productId", service.GetProduct)
				products.PATCH("/:productId", middleware.Authorized, service.UpdateProduct)
				products.DELETE("/:productId", middleware.Authorized, service.DeleteProduct)

			}
			orders := uniqueStore.Group("/orders")
			orders.GET("", service.GetOrders)
		}

	}

	err := router.Run("127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
}
