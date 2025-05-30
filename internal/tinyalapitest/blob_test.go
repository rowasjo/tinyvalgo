package tinyvalapitest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
)

const (
	unknown_blob_sha256_hash  = "61a04a46afa3c518551c887c6c1b2b1e4f25619fad3032c3d5c952849b2ff9db"
	example1_blob             = "I am a little blob."
	example1_blob_sha256_hash = "bfb272e79d30466cf1af7c16739659e8b4e9b85b5075bdb922806c55035497cf"
)

func TestGetBlobInvalidHashReturns400(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	req, err := http.NewRequest(http.MethodGet, "/blobs/invalid-hash", nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusBadRequest)
}

func TestGetUnknownBlobReturns404(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	req, err := http.NewRequest(http.MethodGet, blobUrl(unknown_blob_sha256_hash), nil)
	is.NoErr(err)

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusNotFound)
}

func TestPutBlobWithHashMismatchReturns422(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	body := strings.NewReader("body not matching hash")
	req, err := http.NewRequest(http.MethodPut, blobUrl(unknown_blob_sha256_hash), body)
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/octet-stream")

	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusUnprocessableEntity)
}

func TestPutBlobWithValidHashReturns204(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	rr := putExample1Blob(t, is, app)

	t.Log(rr.Body.String())

	is.Equal(rr.Code, http.StatusNoContent)
}

func TestPutBlobThenGet(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	rr := putExample1Blob(t, is, app)
	is.Equal(rr.Code, http.StatusNoContent)

	req, err := http.NewRequest(http.MethodGet, example1BlobURL, nil)
	is.NoErr(err)

	rr = httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
	is.Equal(rr.Body.String(), example1_blob)
}

func TestPubBlobThenHEAD(t *testing.T) {
	is := is.New(t)
	app := NewTestApp(t)

	rr := putExample1Blob(t, is, app)
	is.Equal(rr.Code, http.StatusNoContent)

	req, err := http.NewRequest(http.MethodHead, example1BlobURL, nil)
	is.NoErr(err)

	rr = httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	is.Equal(rr.Code, http.StatusOK)
	is.Equal(rr.Body.Len(), 0)
	is.Equal(rr.Header().Get("Content-Length"), "19")
}

func putExample1Blob(t *testing.T, is *is.I, handler http.Handler) *httptest.ResponseRecorder {
	t.Helper()

	body := strings.NewReader(example1_blob)
	req, err := http.NewRequest(http.MethodPut, example1BlobURL, body)
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/octet-stream")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func blobUrl(hash string) string {
	return fmt.Sprintf("/blobs/%s", hash)
}

var example1BlobURL = blobUrl(example1_blob_sha256_hash)
