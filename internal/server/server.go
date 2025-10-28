package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	articleUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	authUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"

	articleHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	authHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/router"
)

func StartServer() {
	sessionRepo := session.NewInMemorySession()
	userRepo := user.NewInMemoryUser()
	articleRepo := article.NewInMemoryArticle()

	articleUsecase := articleUsecase.NewArticleUsecase(articleRepo, sessionRepo)
	authUsecase := authUsecase.NewAuthUsecase(userRepo, sessionRepo)

	articleHandler := articleHandler.NewArticleHandler(articleUsecase)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	mux := router.NewRouter(articleHandler, authHandler)
	handler := middleware.CORSMiddleware(mux)

	server := http.Server{
		Addr:         ":8090",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8090")
	server.ListenAndServe()
}
