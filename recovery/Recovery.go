package recovery

import (
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

				log.Println(rl.PanicString(err))
				w.WriteHeader(rl.Status)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
