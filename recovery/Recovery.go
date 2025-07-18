package recovery

import (
	"fmt"
	"log"
	"net/http"

	l "github.com/duartqx/ddgomiddlewares/logger"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rl := l.RequestLogger{
					Method: r.Method,
					Status: http.StatusInternalServerError,
					Path:   r.URL.Path,
				}

				result := fmt.Sprintf(`{"error":"%v"}`, err)

				log.Println(rl.PanicString(result))

				w.WriteHeader(rl.Status)
				w.Write([]byte(result))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
