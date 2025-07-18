package logger

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &ResponseRecorderWriter{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(writer, r)

		log.Println(
			RequestLogger{
				Method: r.Method,
				Result: writer.Result,
				Status: writer.Status,
				Path:   r.URL.Path,
				Since:  time.Since(start),
			},
		)
	})
}
