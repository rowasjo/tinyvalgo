// Package openapidoc provides access to the OpenAPI specification
// for the tinyval service. Clients can import this package to
// retrieve the OpenAPI document programmatically.
package openapidoc

import _ "embed"

//go:embed openapi.yaml
var OpenapiDocument []byte
