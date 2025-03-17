# Tinyval Go

Hash table service for storing and retrieving binary blobs using their SHA-256 hashes as keys.

Note: This project is primarily an experiment in using Go for HTTP service development.

## OpenAPI

Swagger UI is served at `/docs` using the service's [OpenAPI schema](openapidoc/openapi.yaml).

## Development

### Run

```
cd <source-dir>
go run cmd/tinyvalapi/main.go 
```

### Run Tests

```
go test ./...
```
