package apiserver

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
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
	err := validateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.NotFound(w, r)
}

func putBlobHandler(w http.ResponseWriter, r *http.Request) {
	err := validateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.NotFound(w, r)
}

func validateRequest(r *http.Request) error {
	// TODO: parse schema only once, refactor

	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromData(assets.OpenapiYaml)
	if err != nil {
		panic(err)
	}

	router, _ := gorillamux.NewRouter(doc)

	route, pathParams, _ := router.FindRoute(r)

	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
	}
	return openapi3filter.ValidateRequest(ctx, requestValidationInput)
}
