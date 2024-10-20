package tinyvalapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenApiHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/openapi.yaml", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewServer()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected status code %v. Expected %v", status, http.StatusOK)
	}

}
