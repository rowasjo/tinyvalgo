package lib

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

func LoadOpenapiDoc(data []byte) *openapi3.T {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromData(data)
	if err != nil {
		panic(err)
	}
	return doc
}

func OpenAPIValidationMiddlewareFactory(doc *openapi3.T) func(next http.Handler) http.Handler {
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		return makeOpenAPIValidationMiddleware(next, router)
	}
}

func makeOpenAPIValidationMiddleware(next http.Handler, router routers.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, pathParams, err := router.FindRoute(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
		}
		err = openapi3filter.ValidateRequest(r.Context(), requestValidationInput)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
