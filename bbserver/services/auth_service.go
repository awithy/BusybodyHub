package services

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func NewToken(username string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["userid"] = username
	token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString([]byte("secret"))
	return tokenString, err
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err == nil && token.Valid {
		return nil
	} else {
		return errors.New("Invalid token")
	}
}

func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err == nil && token.Valid {
		token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
		tokenString, err := token.SignedString([]byte("secret"))
		return tokenString, err

	} else {
		return "", errors.New("Invalid token")
	}
}
