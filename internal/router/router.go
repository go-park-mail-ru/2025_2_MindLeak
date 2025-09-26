package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
)

func NewRouter(sessions *repository.InMemorySession, users *repository.InMemoryUser) *http.ServeMux {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
	// 	handler.FeedHandler(w, r,)
	// })
	mux.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		handler.RegistrationHandler(w, r, sessions, users)
	})
	mux.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		handler.LoginHandler(w, r, sessions, users)
	})
	mux.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		handler.LogoutHandler(w, r, sessions)
	})

	return mux
}
