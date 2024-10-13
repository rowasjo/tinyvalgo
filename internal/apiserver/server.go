package apiserver

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/lib"
)

const (
	headerContentType = "Content-Type"
	contentTypeHTML   = "text/html"
	contentTypeYAML   = "application/yaml"
)

func ApiServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/openapi.yaml", openapiHandler)
	mux.HandleFunc("/docs", docsHandler)

	doc := lib.OpenapiDoc()
	validation := lib.OpenAPIValidationMiddlewareFactory(doc)

	mux.Handle("GET /blobs/{hash}", validation(http.HandlerFunc(getBlobHandler))) // also matches HEAD
	mux.Handle("PUT /blobs/{hash}", validation(http.HandlerFunc(putBlobHandler)))
	return mux
}

func getBlobHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func putBlobHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
