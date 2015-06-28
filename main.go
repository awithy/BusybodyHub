package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/hosts", HostIndex)
	router.HandleFunc("/hosts/{hostId}", HostDetail)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
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

type Host struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type Hosts []Host
