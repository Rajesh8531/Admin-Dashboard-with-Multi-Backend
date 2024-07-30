package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSize(c *gin.Context) {

	userId := c.Request.Header.Get("id")
	fmt.Println(userId, "USERID")

	storeId := c.Param("storeId")
	fmt.Println(storeId, "storeId")

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "UnAuthorized")
		return
	}

	var size types.Size

	c.BindJSON(&size)

	err := db.CreateSize(size.Name, size.Value, storeId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusConflict, "Error while creating a Store")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func GetSizes(c *gin.Context) {
	storeId := c.Param("storeId")

	sizes := db.GetSizes("storeId = ? ", storeId)

	c.JSON(http.StatusAccepted, sizes)
}

func GetSize(c *gin.Context) {
	storeId := c.Param("storeId")
	sizeId := c.Param("sizeId")

	if sizeId == "new" {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	var size types.Size

	err := db.GetSize(&size, " id = ? AND storeId = ?", sizeId, storeId)
	fmt.Println(size, "SIZE")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, "No size Found")
		return
	}

	c.JSON(http.StatusAccepted, size)
}

func UpdateSize(c *gin.Context) {
	storeId := c.Param("storeId")
	sizeId := c.Param("sizeId")

	var size types.Size
	c.BindJSON(&size)
	fmt.Println("STOREID", storeId)
	fmt.Println("SIZEID", sizeId)
	fmt.Println("NAME", size.Name)
	fmt.Println("VALUE", size.Value)

	err := db.UpdateItem("size", "SET name = ?, value = ? WHERE id = ? AND storeId = ?", size.Name, size.Value, sizeId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Updated")
}

func DeleteSize(c *gin.Context) {
	storeId := c.Param("storeId")
	sizeId := c.Param("sizeId")

	err := db.DeleteItem("size", "WHERE id = ? AND storeId = ? ", sizeId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Deleted")
}
