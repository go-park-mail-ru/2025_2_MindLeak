package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
)

func NewRouter(sessions *repository.InMemorySession, users *repository.InMemoryUser) *http.ServeMux {
	mux := http.NewServeMux()

	//This is a test endpoint just to give user a guest cookie
	mux.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		handler.FeedHandler(w, sessions)
	})

	mux.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		handler.RegistrationHandler(w, r, sessions, users)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handler.LoginHandler(w, r, sessions, users)
	})
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		handler.LogoutHandler(w, r, sessions)
	})

	return mux
}
