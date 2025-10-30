package apiserver

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	articleUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	authUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"

	"github.com/sirupsen/logrus"

	articleHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	authHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/server"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/router"
)

type Server struct {
	config *server.Config
	logger *logrus.Logger
	server http.Server
}

func New(config *server.Config) *Server {
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
		Addr:         config.BindAddr,
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	return &Server{
		config: config,
		logger: logrus.New(),
		server: server,
	}
}

func (s *Server) StartServer() {
	s.logger.Info("starting apiserver")
	s.server.ListenAndServe()
}
