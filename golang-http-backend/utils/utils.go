package utils

import (
	"connection-to-mongo/project/db"
	"connection-to-mongo/project/types"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ResponseJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(payload)
	return err
}

func ResponseError(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), status)
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string, id string) (string, error) {
	envFile, _ := godotenv.Read(".env")
	mySecretKey := []byte(envFile["JWT_SECRET_KEY"])

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

func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	envFile, _ := godotenv.Read(".env")
	mySecretKey := []byte(envFile["JWT_SECRET_KEY"])
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySecretKey), nil
	})

	return claims, err
}

func IsAuthorizedForStore(r *http.Request, storeId primitive.ObjectID) bool {
	var userId = r.Header.Get("id")
	var store types.Store
	db.GetStoreById(&store, bson.M{"storeId": storeId})

	UserId, _ := primitive.ObjectIDFromHex(userId)

	return UserId == store.UserId
}
