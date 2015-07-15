package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	requestpath := r.URL.Path
	if requestpath == "/" {
		requestpath = "/index.html"
	}
	fp := path.Join("bbwebapp/public", requestpath)
	log.Printf("Serving file " + fp)
	http.ServeFile(w, r, fp)
}

func HostIndex(w http.ResponseWriter, r *http.Request) {
	hosts := Hosts{
		Host{Name: "Host 1"},
		Host{Name: "Host 2"},
	}

	//TODO: Refactor
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(hosts); err != nil {
		panic(err)
	}
}

func HostDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hostId := vars["hostId"]
	host := Host{Name: "Host " + hostId}

	//TODO: Refactor
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(host); err != nil {
		panic(err)
	}
}

func HostCreate(w http.ResponseWriter, r *http.Request) {
	var host Host
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &host); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(host); err != nil {
		panic(err)
	}
}

func LoginAccount(w http.ResponseWriter, r *http.Request) {

	var login Login
	var loginResult LoginResult
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
		loginResult = LoginResult{
			true,
			"Login successful",
		}
	} else {
		loginResult = LoginResult{
			false,
			"Invalid username and/or password",
		}
	}

	if err := json.NewEncoder(w).Encode(loginResult); err != nil {
		panic(err)
	}
}
