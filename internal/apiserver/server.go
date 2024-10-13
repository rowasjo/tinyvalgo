package apiserver

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
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

	doc := openapiDoc()
	router := openapiRouter(doc)

	validationMiddleware := func(next http.Handler) http.Handler {
		return makeValidationMiddleware(next, router)
	}

	mux.Handle("GET /blobs/{hash}", validationMiddleware(http.HandlerFunc(getBlobHandler))) // also matches HEAD
	mux.Handle("PUT /blobs/{hash}", validationMiddleware(http.HandlerFunc(putBlobHandler)))
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

func openapiDoc() *openapi3.T {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromData(assets.OpenapiYaml)
	if err != nil {
		panic(err)
	}
	return doc
}

func openapiRouter(doc *openapi3.T) routers.Router {
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		panic(err)
	}
	return router
}

func makeValidationMiddleware(next http.Handler, router routers.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, pathParams, _ := router.FindRoute(r)

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
		}
		err := openapi3filter.ValidateRequest(r.Context(), requestValidationInput)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
