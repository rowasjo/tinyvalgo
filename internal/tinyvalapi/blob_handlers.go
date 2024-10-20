package tinyvalapi

import "net/http"

func getBlobHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func putBlobHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
