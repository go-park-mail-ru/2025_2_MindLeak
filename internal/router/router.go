package router

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/swagger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(articleHandler *article.Handler, authHandler *auth.Handler) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.RecoverMiddleware)

	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.CORSMiddleware)
	//router.Use(middleware.AuthMiddleware) Потом подключить к нужным ручкам

	router.HandleFunc("/feed", articleHandler.Feed).Methods("GET")
	router.HandleFunc("/registration", authHandler.Registration).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("GET")
	router.HandleFunc("/me", authHandler.Me).Methods("GET")
	router.HandleFunc("/swagger", swagger.SwaggerHandler).Methods("GET")

	return router
}
