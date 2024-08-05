package apiserver

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/assets"
)

const (
	headerContentType = "Content-Type"
	contentTypeHTML   = "text/html"
	contentTypeYAML   = "application/yaml"
)

func ApiServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/openapi.yaml", openApiHandler)
	mux.HandleFunc("/docs", docsHandler)

	mux.HandleFunc("GET /blobs/{hash}", getBlobHandler) // also matches HEAD
	mux.HandleFunc("PUT /blobs/{hash}", putBlobHandler)
	return mux
}

func openApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeYAML)
	w.Write(assets.OpenapiYaml)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeHTML)
	w.Write(assets.DocsHtml)
}

func getBlobHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func putBlobHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
