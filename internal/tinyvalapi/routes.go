package tinyvalapi

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/lib"
	"github.com/rowasjo/tinyvalgo/openapidoc"
)

func addRoutes(
	mux *http.ServeMux,
) {
	mux.HandleFunc("/openapi.yaml", openapiHandler)
	mux.HandleFunc("/docs", docsHandler)

	validation := lib.OpenAPIValidationMiddlewareFactory(
		lib.LoadOpenapiDoc(openapidoc.OpenapiDocument))

	mux.Handle("GET /blobs/{hash}", validation(http.HandlerFunc(getBlobHandler))) // also matches HEAD
	mux.Handle("PUT /blobs/{hash}", validation(http.HandlerFunc(putBlobHandler)))
}
