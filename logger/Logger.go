package logger

import (
	"fmt"
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

		defer func() {
			if rec := recover(); rec != nil {
				rl := RequestLogger{
					Method: r.Method,
					Status: http.StatusInternalServerError,
					Path:   r.URL.Path,
					Since:  time.Since(start),
				}

				result := fmt.Sprintf(`{"error":"%v"}`, rec)

				log.Println(rl.PanicString(result))

				writer.WriteHeader(rl.Status)
				writer.Write([]byte(result))

				return
			}

			log.Println(
				RequestLogger{
					Method: r.Method,
					Result: writer.Result,
					Status: writer.Status,
					Path:   r.URL.Path,
					Since:  time.Since(start),
				},
			)

		}()

		next.ServeHTTP(writer, r)
	})
}
