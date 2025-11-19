package middleware

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"net/http"
)

// future auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := cookies.GetCookie(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		next.ServeHTTP(w, r)
	})
}
