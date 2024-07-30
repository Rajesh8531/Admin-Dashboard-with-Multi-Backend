package middleware

import (
	"admin-dashboard/backend/golan-gin/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {

	authorizationString := c.Request.Header.Get("Authorization")

	if authorizationString == "" {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	tokenString := strings.Split(authorizationString, ` `)[1]
	claims, _ := utils.DecodeJWT(tokenString)

	var id string = claims["id"].(string)

	c.Request.Header.Add("id", id)

	c.Next()

}
