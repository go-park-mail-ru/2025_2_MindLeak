package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error(r.Context(), "PANIC: %v\n%s", rec, string(debug.Stack()), nil)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
