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

func validationMiddleware(next http.Handler) http.Handler {

	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromData(assets.OpenapiYaml)
	if err != nil {
		panic(err)
	}

	router, _ := gorillamux.NewRouter(doc)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, pathParams, _ := router.FindRoute(r)

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
		}
		err := openapi3filter.ValidateRequest(ctx, requestValidationInput)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
