package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

func AuthToken(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err == nil && token.Valid {
			inner.ServeHTTP(w, r)
			log.Printf("Auth success")
		} else {
			log.Printf("Auth failed")
		}
	})
}
