package tinyvalapi

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/lib"
)

func NewServer(repo lib.Repository) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, repo)
	var handler http.Handler = mux
	return handler
}
