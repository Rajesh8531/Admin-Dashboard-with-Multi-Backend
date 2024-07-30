package middleware

import (
	"connection-to-mongo/project/utils"
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		requestString := r.Header.Get("Authorization")
		if requestString == "" {
			utils.ResponseError(w, http.StatusUnauthorized, fmt.Errorf("aunauthorized"))
			return
		}

		tokenString := strings.Split(requestString, ` `)[1]

		claims, _ := utils.DecodeJWT(tokenString)

		var email string = claims["email"].(string)
		var id string = claims["id"].(string)
		r.Header.Add("name", email)
		r.Header.Add("id", id)
		next.ServeHTTP(w, r)
	})
}
