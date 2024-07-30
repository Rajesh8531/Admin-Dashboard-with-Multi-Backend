package service

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"admin-dashboard/backend/golan-gin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SignUp(c *gin.Context) {
	var user types.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "No data found"})
		return
	}

	var existingUser types.User

	_ = db.GetUser(&existingUser, "WHERE email = ?", user.Email)

	if existingUser.Email != "" {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "user Already Existing"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Error in Password Generation"})
		return
	}

	id, err := db.CreateUser(user.Name, user.Email, hashedPassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "user Already Existing"})
		return
	}

	envFile, err := godotenv.Read(".env")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "user Already Existing"})
		return
	}
	secret := envFile["JWT_SECRET_KEY"]
	tokenString, err := utils.GenerateJWT(user.Email, id.String(), secret)
	if err != nil {
		fmt.Println(err)
		return
	}
	// c.SetCookie("authentication", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, gin.H{"token": tokenString, "name": user.Name, "email": user.Email, "id": id})
}

func SignIn(c *gin.Context) {

	var user types.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "No data given"})
		return
	}

	var existingUser types.User

	if err := db.GetUser(&existingUser, "WHERE email = ? ", user.Email); err != nil {
		fmt.Print(err)
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "Error while fetching the user"})
		return
	}

	if existingUser.Email == "" {
		c.AbortWithError(http.StatusConflict, fmt.Errorf("no user Found"))
		return
	}

	isPasswordCorrect := utils.CheckPassword(user.Password, existingUser.Password)

	if !isPasswordCorrect {
		c.AbortWithError(http.StatusNotAcceptable, fmt.Errorf("invalid Credentials"))
		return
	}

	envFile, err := godotenv.Read(".env")

	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("no secrets found"))
		return
	}

	tokenString, err := utils.GenerateJWT(existingUser.Email, existingUser.ID, envFile["JWT_SECRET_KEY"])
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("error while generating JWT"))
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": tokenString, "name": existingUser.Name, "email": existingUser.Email, "id": existingUser.ID})
}
