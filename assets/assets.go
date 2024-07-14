package assets

import _ "embed"

//go:embed openapi.yaml
var OpenapiYaml []byte

//go:embed docs.html
var DocsHtml []byte
