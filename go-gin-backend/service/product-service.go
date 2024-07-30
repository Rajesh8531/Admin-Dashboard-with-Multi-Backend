package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	userId := c.Request.Header.Get("id")

	storeId := c.Param("storeId")

	isAuthorized := utils.IsAuthorized(userId, storeId)

	if !isAuthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "UnAuthorized")
		return
	}

	var product types.ProductWithImageUrl

	c.BindJSON(&product)

	id, err := db.CreateProduct(product.Name, product.CategoryId, storeId, product.IsFeatured, product.IsArchived, product.SizeId, product.ColorId, product.Price)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusConflict, "Error while creating a Category")
		return
	}

	for _, url := range product.ImageUrl {
		err := db.CreateImage(url, id.String())
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusConflict, "Error while creating a Image")
			return
		}
	}

	c.JSON(http.StatusAccepted, "Product Created Successfully")
}

func GetProducts(c *gin.Context) {
	storeId := c.Param("storeId")

	fullProducts := db.GetFullProducts(" storeId = ? ", storeId)
	fmt.Println(fullProducts)

	c.JSON(http.StatusAccepted, fullProducts)
}

func GetProduct(c *gin.Context) {
	storeId := c.Param("storeId")
	productId := c.Param("productId")
	if productId == "new" {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	var product types.FullProduct

	err := db.GetProduct(&product, "id = ? AND storeId = ?", productId, storeId)

	var category types.Category
	var color types.Color
	var size types.Size
	var images []types.Image

	db.GetCategory(&category, " id = ? ", product.CategoryId)
	db.GetColor(&color, " id = ? ", product.ColorId)
	images = db.GetImages("productId = ?", product.ID)
	db.GetSize(&size, "id = ?", product.SizeId)

	product.Category = category
	product.Color = color
	product.Image = images
	product.Size = size

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, "No category Found")
		return
	}

	c.JSON(http.StatusAccepted, product)
}

func UpdateProduct(c *gin.Context) {
	storeId := c.Param("storeId")
	productId := c.Param("productId")

	var product types.ProductWithImageUrl
	c.BindJSON(&product)

	fmt.Println(productId, "PRODUCT ID")
	err := db.DeleteItem("image", "WHERE productId = ? ", productId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "ERROR WHILE UPDATING IMAGE")
		return
	}

	err = db.UpdateItem(`product`, `
		SET name = ?, 
			price = ?,
			categoryId = ?,
			isArchived= ?,
			isFeatured = ?,
			colorId = ?, 
			sizeId = ?, 
			storeId = ?
			WHERE id = ? `,

		product.Name,
		product.Price,
		product.CategoryId,
		product.IsArchived,
		product.IsFeatured,
		product.ColorId,
		product.SizeId,
		storeId,
		productId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "ERROR WHILE UPDATING PRODUCT")
		return
	}

	for _, url := range product.ImageUrl {
		err = db.CreateImage(url, productId)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}

	c.JSON(http.StatusAccepted, "Successfully Updated")
}

func DeleteProduct(c *gin.Context) {
	storeId := c.Param("storeId")
	productId := c.Param("productId")
	err := db.DeleteItem("image", "WHERE productId = ? ", productId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	err = db.DeleteItem("product", "WHERE id = ? AND storeId = ? ", productId, storeId)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusAccepted, "Successfully Deleted")
}
