package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStore(c *gin.Context) {
	userId := c.Request.Header.Get("id")

	var store types.Store

	c.BindJSON(&store)

	id, err := db.CreateStore(store.Name, userId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, "error while creating store")
		return
	}

	c.JSON(http.StatusAccepted, id)
}

func GetStore(c *gin.Context) {
	storeId := c.Param("storeId")
	var store types.FullStoreType

	err := db.GetStore(&store, " id = ? ", storeId)
	if err != nil {
		fmt.Println("Something Went wrong CREATE STORE CONTROLLER")
		return
	}
	store.Sizes = db.GetSizes(" storeId = ?", storeId)
	store.Billboards, _ = db.GetBillboards(" storeId = ?", storeId)
	store.Colors = db.GetColors(" storeId = ?", storeId)
	store.Categories = db.GetCategories(" storeId = ?", storeId)

	c.JSON(http.StatusAccepted, store)
}

func GetStores(c *gin.Context) {
	userId := c.Query("userId")

	stores := db.GetFullStores(" userId = ? ", userId)

	c.JSON(http.StatusAccepted, stores)
}

func UpdateStore(c *gin.Context) {
	userId := c.Request.Header.Get("id")
	storeId := c.Param("storeId")

	var store types.Store

	c.BindJSON(&store)

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Unauthorized")
		return
	}

	err := db.UpdateItem("store", "SET name = ? WHERE id = ? AND userId = ?", store.Name, storeId, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error while Updating Store")
		return
	}

	c.JSON(http.StatusAccepted, store)
}

func DeleteStore(c *gin.Context) {
	userId := c.Request.Header.Get("id")
	storeId := c.Param("storeId")

	var store types.Store

	c.BindJSON(&store)

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Unauthorized")
		return
	}

	err := db.DeleteItem("store", "WHERE id = ? AND userId = ? ", storeId, userId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Deleted")
}
