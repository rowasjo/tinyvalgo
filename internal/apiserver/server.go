package apiserver

import (
	"log"
	"net/http"

	"github.com/rowasjo/tinyvalgo/assets"
)

const (
	headerContentType = "Content-Type"
	contentTypeHTML   = "text/html"
	contentTypeYAML   = "application/yaml"
	swaggerUiHtml     = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="description" content="SwaggerUI" />
	<title>Tinyval - Swagger UI</title>
	<link type="text/css" rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@4.5.0/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@4.5.0/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {{
	window.ui = SwaggerUIBundle({{
		url: '/openapi.yaml',
		dom_id: '#swagger-ui',
		layout: "BaseLayout",
		deepLinking: true,
		displayRequestDuration: true,
		showExtensions: true,
		showCommonExtensions: true,
		presets: [
		SwaggerUIBundle.presets.apis,
		SwaggerUIBundle.SwaggerUIStandalonePreset
		],
	}});
	}};
</script>
</body>
</html>`
)

func ApiServer() {
	http.HandleFunc("/openapi.yaml", openApiHandler)
	http.HandleFunc("/docs", docsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func openApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeYAML)
	w.Write([]byte(assets.OpenapiYaml))
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(headerContentType, contentTypeHTML)
	w.Write([]byte(swaggerUiHtml))
}

// func readOpenApiYaml() string {
// 	exePath, err := os.Executable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rootPath := filepath.Dir(filepath.Dir(exePath))

// 	filePath := filepath.Join(rootPath, "assets", "openapi.yaml")

// 	content, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return string(content)
// }
