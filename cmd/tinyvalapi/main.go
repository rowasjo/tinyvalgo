package main

import (
	"log"
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

func main() {
	handler := tinyvalapi.NewServer()

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
