package apiserver

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/assets"

	_ "embed"
)

//go:embed swaggerui.html
var swaggeruiHTML []byte

func openapiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeYAML)
	w.Write(assets.OpenapiYAML)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeHTML)
	w.Write(swaggeruiHTML)
}
