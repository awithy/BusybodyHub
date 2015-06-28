package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
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
