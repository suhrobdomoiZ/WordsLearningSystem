package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter

	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		respWr := responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		start := time.Now()

		next.ServeHTTP(&respWr, req)

		elapsed := time.Since(start)

		logger.Info("http request",
			slog.String("method", req.Method),
			slog.String("path", req.URL.Path),
			slog.String("status", http.StatusText(respWr.statusCode)),
			slog.Duration("duration", elapsed),
		)
	})
}
