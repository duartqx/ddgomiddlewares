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
				rl := l.NewRequestLoggerBuilder().
					SetMethod(r.Method).
					SetStatus(http.StatusInternalServerError).
					SetPath(r.URL.Path)

				log.Println(rl.PanicString(err))
				w.WriteHeader(rl.GetStatus())
			}
		}()
		next.ServeHTTP(w, r)
	})
}
