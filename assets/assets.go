package assets

import _ "embed"

//go:embed openapi.yaml
var OpenapiYAML []byte

//go:embed docs.html
var DocsHTML []byte
