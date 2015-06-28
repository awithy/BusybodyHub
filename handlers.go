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
	json.NewEncoder(w).Encode(hosts)
}

func HostDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hostId := vars["hostId"]
	host := Host{Name: "Host " + hostId}
	json.NewEncoder(w).Encode(host)
}
