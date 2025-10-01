package router

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/feed"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/login"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/logout"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/registration"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"

	handler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/me"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

func NewRouter(sessions *session.InMemorySession, users *user.InMemoryUser, articles *article.InMemoryArticle) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/feed", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			feed.FeedHandler(w, r, sessions, articles)
		},
	)))

	mux.Handle("/registration", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			registration.RegistrationHandler(w, r, sessions, users)
		},
	)))

	mux.Handle("/login", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			login.LoginHandler(w, r, sessions, users)
		},
	)))

	mux.Handle("/logout", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logout.LogoutHandler(w, r, sessions)
		},
	)))

	mux.Handle("/me", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler.MeHandler(w, r, sessions, users)
		},
	)))

	return mux
}
