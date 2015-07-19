package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/awithy/busybodyhub/bbserver/services"
)

func AuthToken(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		err := services.VerifyToken(tokenString)
		if err == nil {
			inner.ServeHTTP(w, r)
			log.Printf("Auth success")
		} else {
			log.Printf("Auth failed")
		}
	})
}
