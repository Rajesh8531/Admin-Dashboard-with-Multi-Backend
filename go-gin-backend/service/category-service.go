package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	userId := c.Request.Header.Get("id")

	storeId := c.Param("storeId")

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "UnAuthorized")
		return
	}

	var category types.Category

	c.BindJSON(&category)

	err := db.CreateCategory(category.Name, storeId, category.BillboardId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusConflict, "Error while creating a Category")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success"})
}

func GetCategories(c *gin.Context) {
	storeId := c.Param("storeId")

	categories := db.GetFullCategories(" storeId = ? ", storeId)

	c.JSON(http.StatusAccepted, categories)
}

func GetCategory(c *gin.Context) {
	storeId := c.Param("storeId")
	categoryId := c.Param("categoryId")

	if categoryId == "new" {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	var category types.FullCategory

	var cat types.Category
	err := db.GetCategory(&cat, " id = ? ", categoryId)

	category.CreatedAt = cat.CreatedAt
	category.UpdatedAt = cat.UpdatedAt
	category.ID = cat.ID
	category.Name = cat.Name
	category.StoreId = cat.StoreId
	category.BillboardId = cat.BillboardId

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, "No category Found")
		return
	}

	err = db.GetBillboard(&category.Billboard, " id = ? AND storeId = ? ", cat.BillboardId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error while fetching Billboard")
		return
	}

	category.Products = db.GetProducts(" categoryId = ? ", categoryId)

	c.JSON(http.StatusAccepted, category)
}

func UpdateCategory(c *gin.Context) {
	storeId := c.Param("storeId")
	categoryId := c.Param("categoryId")

	var category types.Category
	c.BindJSON(&category)

	err := db.UpdateItem("category", "SET name = ?, billboardId = ? WHERE id = ? AND storeId = ?", category.Name, category.BillboardId, categoryId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Updated")
}

func DeleteCategory(c *gin.Context) {
	storeId := c.Param("storeId")
	categoryId := c.Param("categoryId")

	err := db.DeleteItem("category", "WHERE id = ? AND storeId = ? ", categoryId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Deleted")
}
