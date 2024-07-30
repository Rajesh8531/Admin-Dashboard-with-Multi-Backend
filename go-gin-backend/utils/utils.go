package utils

import (
	"admin-dashboard/backend/golan-gin/db"
	"admin-dashboard/backend/golan-gin/types"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	envFile, _ := godotenv.Read(".env")
	mySecretKey := []byte(envFile["JWT_SECRET_KEY"])
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySecretKey), nil
	})

	return claims, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateJWT(email string, id string, secret string) (string, error) {
	mySecretKey := []byte(secret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["id"] = id

	tokenString, err := token.SignedString(mySecretKey)

	if err != nil {
		fmt.Println("Something went wrong in generating Token string")
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(userId string, storeId string) bool {
	var store types.FullStoreType
	_ = db.GetStore(&store, " id = ? AND userId = ?", storeId, userId)
	return store.ID != ""
}
