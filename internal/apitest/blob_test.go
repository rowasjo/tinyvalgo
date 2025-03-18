package apitest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
)

const (
	unknown_blob_sha256_hash = "61a04a46afa3c518551c887c6c1b2b1e4f25619fad3032c3d5c952849b2ff9db"
)

func TestGetBlobInvalidHashReturns400(t *testing.T) {
	is := is.New(t)
	handler := NewTestServer(t)

	req, err := http.NewRequest("GET", "/blobs/invalid-hash", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusBadRequest)
}

func TestGetUnknownBlobReturns404(t *testing.T) {
	is := is.New(t)
	handler := NewTestServer(t)

	req, err := http.NewRequest("GET", blobUrl(unknown_blob_sha256_hash), nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusNotFound)
}

func TestPutBlobWithInvalidBodyReturns422(t *testing.T) {
	is := is.New(t)
	handler := NewTestServer(t)

	body := strings.NewReader("body not matching hash")
	req, err := http.NewRequest("PUT", blobUrl(unknown_blob_sha256_hash), body)
	req.Header.Set("Content-Type", "application/octet-stream")
	is.NoErr(err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusUnprocessableEntity)
}

func blobUrl(hash string) string {
	return fmt.Sprintf("/blobs/%s", hash)
}
