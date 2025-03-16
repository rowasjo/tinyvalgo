package apitest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

func TestGetOpenApiYAML(t *testing.T) {
	is := is.New(t)
	handler := tinyvalapi.NewServer()

	req, err := http.NewRequest("GET", "/openapi.yaml", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
}

func TestGetSwaggerDocs(t *testing.T) {
	is := is.New(t)
	handler := tinyvalapi.NewServer()

	req, err := http.NewRequest("GET", "/docs", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
}
