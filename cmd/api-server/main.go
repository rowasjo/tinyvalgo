package main

import (
	"log"
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/apiserver"
)

func main() {
	mux := apiserver.ApiServer()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
