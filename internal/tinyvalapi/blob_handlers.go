package tinyvalapi

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/lib"
)

func getBlobHandler(repo lib.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hash := getHashPathParam(r)
		if !repo.Exists(ctx, hash) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

func putBlobHandler(repo lib.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hash := getHashPathParam(r)

		err := repo.Put(ctx, hash, r.Body)
		if err != nil {

		}

		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

func getHashPathParam(r *http.Request) string {
	hash := r.PathValue("hash")
	if hash == "" {
		// OpenAPI validation ensures hash path parameter is a valid SHA-256 hash
		panic("missing hash path parameter")
	}
	return hash
}
