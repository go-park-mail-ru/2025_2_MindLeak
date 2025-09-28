package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
)

func NewRouter(sessions *repository.InMemorySession, users *repository.InMemoryUser, articles *repository.InMemoryArticle) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/feed", middleware.CORSMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handler.FeedHandler(w, r, sessions, articles)
	}))
	mux.HandleFunc("/registration", middleware.CORSMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handler.RegistrationHandler(w, r, sessions, users)
	}))
	mux.HandleFunc("/login", middleware.CORSMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handler.LoginHandler(w, r, sessions, users)
	}))
	mux.HandleFunc("/logout", middleware.CORSMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handler.LogoutHandler(w, r, sessions)
	}))

	return mux
}
