package timeout

import (
	"context"
	"net/http"
	"time"

	i "github.com/duartqx/ddgomiddlewares/interfaces"
)

func TimeoutMiddleware(timeout time.Duration, err error) i.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx, cancel := context.WithTimeout(r.Context(), time.Second*timeout)
			defer cancel()

			done := make(chan bool)
			panicChan := make(chan any, 1)

			go func() {
				defer func() {
					if p := recover(); p != nil {
						panicChan <- p
					}

					close(done)
				}()

				next.ServeHTTP(w, r.WithContext(ctx))
			}()

			select {
			case <-ctx.Done():
				w.WriteHeader(http.StatusGatewayTimeout)
				panic(err)
			case p := <-panicChan:
				panic(p)
			case <-done:
			}
		})
	}
}
