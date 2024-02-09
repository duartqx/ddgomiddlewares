package trailling

import (
	"net/http"
	"strings"
)

func TrailingSlashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/") && r.URL.Path != "" {
			r.URL.Path += "/"
		}
		next.ServeHTTP(w, r)
	})
}
