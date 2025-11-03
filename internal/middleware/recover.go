package middleware

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(r.Context(), "panic recovered: %v", err)
				debug.PrintStack()
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
