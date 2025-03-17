package apitest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestGetOpenApiYAML(t *testing.T) {
	is := is.New(t)
	handler := NewTestServer(t)

	req, err := http.NewRequest("GET", "/openapi.yaml", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
}

func TestGetSwaggerDocs(t *testing.T) {
	is := is.New(t)
	handler := NewTestServer(t)

	req, err := http.NewRequest("GET", "/docs", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
}
