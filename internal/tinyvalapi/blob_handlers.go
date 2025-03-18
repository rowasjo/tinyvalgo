package tinyvalapi

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/rowasjo/tinyvalgo/internal/lib"
)

func getBlobHandler(repo lib.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hash := getHashPathParam(r)

		reader, size, err := repo.Get(ctx, hash)
		if err == lib.ErrNotFound {
			http.Error(w, "not found", http.StatusNotFound)
			return
		} else if err != nil {
			slog.Info("error fetching blob", slog.String("err", err.Error()))
			http.Error(w, "internal error fetching blob", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Cache-Control", "max-age=31536000, immutable")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.FormatInt(size, 10))

		http.ServeContent(w, r, "", time.Time{}, reader)
	}
}

func putBlobHandler(repo lib.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hash := getHashPathParam(r)

		if err := repo.Put(ctx, hash, r.Body); err != nil {
			var hmErr *lib.HashMismatchError
			if ok := errors.As(err, &hmErr); ok {
				http.Error(w, hmErr.Error(), http.StatusUnprocessableEntity)
			} else {
				http.Error(w, "failed to store blob", http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
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
