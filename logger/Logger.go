package logger

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/duartqx/ddgomiddlewares/interfaces"
	"github.com/google/uuid"
)

const X_REQUEST_ID string = "X-Request-Id"

func ternary(condition bool, a, b int) int {
	if condition {
		return a
	}
	return b
}

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
					Status: ternary(
						writer.Status > http.StatusInternalServerError,
						writer.Status,
						http.StatusInternalServerError,
					),
					Path:  r.URL.Path,
					Since: time.Since(start),
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

func SLoggerMiddleware(service string, slogger *slog.Logger) interfaces.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			writer := &ResponseRecorderWriter{
				Id:             uuid.New(),
				ResponseWriter: w,
				Status:         http.StatusOK,
			}

			ctx := context.WithValue(r.Context(), X_REQUEST_ID, writer.Id.String())

			writer.Header().Add(X_REQUEST_ID, writer.Id.String())

			defer func() {
				rl := NewRequestSLogger(writer.Id).WithMethod(r.Method).WithPath(r.URL.Path).WithHost(r.URL.Host)

				if rec := recover(); rec != nil {

					result := fmt.Sprintf(`{"error":"%v"}`, rec)

					slogger.Error(
						service,
						rl.
							WithSince(time.Since(start)).
							WithResult(result).
							WithStatus(
								ternary(
									writer.Status > http.StatusInternalServerError,
									writer.Status,
									http.StatusInternalServerError,
								),
							).
							Slog()...,
					)

					writer.Header().Set("Content-Type", "application/json")
					writer.WriteHeader(rl.status)
					writer.Write([]byte(result))

					return
				}

				slogger.Info(
					service,
					rl.
						WithSince(time.Since(start)).
						WithResult(writer.Result).
						WithStatus(writer.Status).
						Slog()...,
				)

			}()

			next.ServeHTTP(writer, r.WithContext(ctx))
		})
	}
}
