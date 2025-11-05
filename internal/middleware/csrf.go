package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
	"time"
)

const (
	csrfCookieName = "csrf_token"
	csrfHeaderName = "X-CSRF-Token"
)

func generateCSRFToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		method := strings.ToUpper(r.Method)

		if method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions {
			if _, err := r.Cookie(csrfCookieName); err != nil {
				token := generateCSRFToken()
				http.SetCookie(w, &http.Cookie{
					Name:     csrfCookieName,
					Value:    token,
					Path:     "/",
					Secure:   true,
					HttpOnly: true,
					SameSite: http.SameSiteNoneMode,
					Expires:  time.Now().Add(1 * time.Hour),
				})
			}
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie(csrfCookieName)
		if err != nil {
			http.Error(w, "missing csrf cookie", http.StatusForbidden)
			return
		}

		header := r.Header.Get(csrfHeaderName)
		if header == "" {
			http.Error(w, "missing csrf header", http.StatusForbidden)
			return
		}

		if cookie.Value != header {
			http.Error(w, "invalid csrf token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
