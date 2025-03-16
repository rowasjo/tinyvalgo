package apitest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

const (
	unknown_blob_sha256_hash = "61a04a46afa3c518551c887c6c1b2b1e4f25619fad3032c3d5c952849b2ff9db"
)

func TestGetBlobInvalidHashReturns400(t *testing.T) {
	is := is.New(t)
	handler := tinyvalapi.NewServer()

	req, err := http.NewRequest("GET", "/blobs/invalid-hash", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusBadRequest)
}

func TestGetUnknownBlobReturns404(t *testing.T) {
	is := is.New(t)
	handler := tinyvalapi.NewServer()

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("/blobs/%s", unknown_blob_sha256_hash),
		nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusNotFound)
}
