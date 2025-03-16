package lib

import (
	"log/slog"
	"net/http"
)

// OpenTelemetry semantic conventions
const (
	attribute_http_request_method       = "http.request.method"
	attribute_http_response_status_code = "http.response.status.code"
	attribute_url_path                  = "url.path"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		slog.Info("request",
			slog.String(attribute_http_request_method, r.Method),
			slog.String(attribute_url_path, r.URL.Path),
		)
		next.ServeHTTP(lrw, r)
		slog.Info("response",
			slog.String(attribute_http_request_method, r.Method),
			slog.String(attribute_url_path, r.URL.Path),
			slog.Int(attribute_http_response_status_code, lrw.statusCode),
		)
	})
}
