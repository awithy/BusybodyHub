package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/awithy/busybodyhub/bbserver/models"
)

func LoginAccount(w http.ResponseWriter, r *http.Request) {

	var login models.Login
	var loginResult models.LoginResult
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &login); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(200)

	if strings.ToLower(login.Username) == "admin" && login.Password == "busybody" {

		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims["userid"] = login.Username
		token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
		tokenString, err := token.SignedString([]byte("secret"))

		if err != nil {
			panic("Error creating token string")
		}

		loginResult = models.LoginResult{
			true,
			"Login successful",
			tokenString,
		}

	} else {
		loginResult = models.LoginResult{
			false,
			"Invalid username and/or password",
			"",
		}
	}

	if err := json.NewEncoder(w).Encode(loginResult); err != nil {
		panic(err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("Logged out")
}
