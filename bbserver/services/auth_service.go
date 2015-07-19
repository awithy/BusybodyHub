package services

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

func Verify(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err == nil && token.Valid {
		return nil
	} else {
		return errors.New("Invalid token")
	}
}
