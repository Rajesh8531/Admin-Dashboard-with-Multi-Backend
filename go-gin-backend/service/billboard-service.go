package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBillboard(c *gin.Context) {

	userId := c.Request.Header.Get("id")

	storeId := c.Param("storeId")

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "UnAuthorized")
		return
	}

	var billboard types.Billboard

	c.BindJSON(&billboard)

	err := db.CreateBillboard(billboard.Label, storeId, billboard.ImageUrl)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, "Error while creating a billboard")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func GetBillboards(c *gin.Context) {
	storeId := c.Param("storeId")

	billboards, _ := db.GetBillboards(" storeId = ? ", storeId)
	fmt.Println(billboards)

	c.JSON(http.StatusAccepted, billboards)
}

func GetBillboard(c *gin.Context) {
	storeId := c.Param("storeId")
	billboardId := c.Param("billboardId")
	fmt.Println(billboardId, "BILLBOARD ID")

	if billboardId == "new" {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	var billboard types.Billboard
	err := db.GetBillboard(&billboard, " id = ? AND storeId = ? ", billboardId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, billboard)
}

func UpdateBillboard(c *gin.Context) {
	storeId := c.Param("storeId")
	billboardId := c.Param("billboardId")

	var billboard types.Billboard
	c.BindJSON(&billboard)

	err := db.UpdateItem("billboard", "SET label = ?, imageUrl = ? WHERE id = ? AND storeId = ?", billboard.Label, billboard.ImageUrl, billboardId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Updated")
}

func DeleteBillboard(c *gin.Context) {
	storeId := c.Param("storeId")
	billboardId := c.Param("billboardId")

	err := db.DeleteItem("billboard", "WHERE id = ? AND storeId = ? ", billboardId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Deleted")
}
