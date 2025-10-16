package router

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/swagger"
	to_delete_прошлый_код__удалить_как_перепишем_2 "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/to_delete(прошлый код, удалить как перепишем)"

	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/feed"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

func NewRouter(sessions session.SessionRepository, users user.UserRepository, articles article.ArticleRepository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/feed", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			feed.FeedHandler(w, r, sessions, articles)
		},
	)))

	mux.Handle("/registration", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			to_delete_прошлый_код__удалить_как_перепишем_2.RegistrationHandler(w, r, sessions, users)
		},
	)))

	mux.Handle("/login", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			to_delete_прошлый_код__удалить_как_перепишем_2.LoginHandler(w, r, sessions, users)
		},
	)))

	mux.Handle("/logout", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			to_delete_прошлый_код__удалить_как_перепишем_2.LogoutHandler(w, r, sessions)
		},
	)))

	mux.Handle("/me", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			to_delete_прошлый_код__удалить_как_перепишем_2.MeHandler(w, r, sessions, users)
		},
	)))

	mux.Handle("/swagger/", middleware.CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			swagger.SwaggerHandler(w, r)
		},
	)))

	return mux
}
