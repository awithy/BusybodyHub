package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

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
		loginResult = models.LoginResult{
			true,
			"Login successful",
		}
	} else {
		loginResult = models.LoginResult{
			false,
			"Invalid username and/or password",
		}
	}

	if err := json.NewEncoder(w).Encode(loginResult); err != nil {
		panic(err)
	}
}
