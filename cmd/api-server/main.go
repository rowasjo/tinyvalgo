package main

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/internal/apiserver"
)

func main() {
	http.HandleFunc("/", apiserver.HelloServer)
	http.ListenAndServe(":8080", nil)
}
