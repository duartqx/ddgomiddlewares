package cache

import (
	"net/http"
	"time"

	i "github.com/duartqx/ddgomiddlewares/interfaces"
)

func CacheMiddleware(maxAge string, delta time.Duration) i.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", "public, max-age="+maxAge)
			w.Header().Set("Expires", time.Now().Add(delta).Format(http.TimeFormat))

			next.ServeHTTP(w, r)
		})
	}
}
