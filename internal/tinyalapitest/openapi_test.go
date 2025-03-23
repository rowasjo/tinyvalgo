package tinyvalapitest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestGetOpenApiYAML(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	req, err := http.NewRequest(http.MethodGet, "/openapi.yaml", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
}

func TestGetSwaggerDocs(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	req, err := http.NewRequest(http.MethodGet, "/docs", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
}
