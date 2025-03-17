package tinyvalapi

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/lib"
	"github.com/rowasjo/tinyvalgo/openapidoc"
)

func addRoutes(
	mux *http.ServeMux,
	repo lib.Repository,
) {
	coreMiddlewares := lib.LoggingMiddleware

	mux.Handle("/openapi.yaml", coreMiddlewares(http.HandlerFunc(openapiHandler)))
	mux.Handle("/docs", coreMiddlewares(http.HandlerFunc(docsHandler)))

	validation := lib.OpenAPIValidationMiddlewareFactory(
		lib.LoadOpenapiDoc(openapidoc.OpenapiDocument))

	blobMiddlewares := func(h http.Handler) http.Handler {
		return coreMiddlewares(validation(h))
	}

	mux.Handle("GET /blobs/{hash}", blobMiddlewares(http.HandlerFunc(getBlobHandler(repo)))) // also matches HEAD
	mux.Handle("PUT /blobs/{hash}", blobMiddlewares(http.HandlerFunc(putBlobHandler(repo))))
}

type Middleware func(http.Handler) http.Handler
