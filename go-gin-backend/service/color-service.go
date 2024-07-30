package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateColor(c *gin.Context) {

	userId := c.Request.Header.Get("id")

	storeId := c.Param("storeId")

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "UnAuthorized")
		return
	}

	var color types.Color

	c.BindJSON(&color)

	err := db.CreateColor(color.Name, storeId, color.Value)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusConflict, "Error while creating a Color")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func GetColors(c *gin.Context) {
	storeId := c.Param("storeId")

	colors := db.GetColors("storeId = ? ", storeId)

	c.JSON(http.StatusAccepted, colors)
}

func GetColor(c *gin.Context) {
	storeId := c.Param("storeId")
	colorId := c.Param("colorId")
	if colorId == "new" {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	var color types.Color

	err := db.GetColor(&color, " id = ? AND storeId = ?", colorId, storeId)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, "No color Found")
		return
	}

	c.JSON(http.StatusAccepted, color)
}

func UpdateColor(c *gin.Context) {
	storeId := c.Param("storeId")
	colorId := c.Param("colorId")

	var color types.Color
	c.BindJSON(&color)

	err := db.UpdateItem("color", "SET name = ?, value = ? WHERE id = ? AND storeId = ?", color.Name, color.Value, colorId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Updated")
}

func DeleteColor(c *gin.Context) {
	storeId := c.Param("storeId")
	colorId := c.Param("colorId")

	err := db.DeleteItem("color", "WHERE id = ? AND storeId = ? ", colorId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Deleted")
}
