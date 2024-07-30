package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	storeId := c.Param("storeId")

	orders := db.GetOrders("storeId = ? ", storeId)

	c.JSON(http.StatusAccepted, orders)
}
