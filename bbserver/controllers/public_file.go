package controllers

import (
	"log"
	"net/http"
	"path"
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
