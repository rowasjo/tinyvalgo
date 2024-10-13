package apiserver

import (
	"net/http"

	"github.com/rowasjo/tinyvalgo/openapidoc"

	_ "embed"
)

//go:embed swaggerui.html
var swaggeruiHTML []byte

const (
	headerContentType = "Content-Type"
	contentTypeHTML   = "text/html"
	contentTypeYAML   = "application/yaml"
)

func openapiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeYAML)
	w.Write(openapidoc.OpenapiDocument)
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeHTML)
	w.Write(swaggeruiHTML)
}
