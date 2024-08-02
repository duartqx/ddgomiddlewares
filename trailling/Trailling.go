package trailling

import (
	"net/http"
	"strings"
)

func TrailingSlashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			if path, found := strings.CutSuffix(r.URL.Path, "/"); found {
				r.URL.Path = path
			}
		}
		next.ServeHTTP(w, r)
	})
}
