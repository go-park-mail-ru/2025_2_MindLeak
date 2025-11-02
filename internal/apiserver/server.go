package apiserver

import (
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/minio"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/postgres"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/redis"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	articleUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	authUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"
	profileUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/usecase"

	articleHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	authHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	profileHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/server"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/router"
)

type Server struct {
	config *server.Config
	server http.Server
}

func New(config *server.Config) (*Server, error) {
	logger.Info(nil, "Logger initialized", nil)
	logger.Info(nil, "Server started initializing", nil)

	//INITIALIZE POSTGRES
	PGConfig := postgres.NewPostgresConfig()
	DB, err := PGConfig.PGconnect()
	if err != nil {
		logger.Error(nil, "Error initializing DB connection", nil)
		return nil, err
	}

	logger.Info(nil, "Postgres connection initialized", nil)

	//INITIALIZE REDIS
	RedisConfig := redis.NewRedisConfig()
	RedisConn, err := RedisConfig.RedisConnect()
	if err != nil {
		logger.Error(nil, "Error initializing Redis connection", nil)
		return nil, err
	}

	logger.Info(nil, "Redis connection initialized", nil)

	//INITIALIZE MINIO
	_ = minio.NewMinioConfig()

	sessionRepo := session.NewRedisSessionManager(RedisConn)
	userRepo := user.NewPostgresUser(DB)
	articleRepo := article.NewArticleRepo(DB)
	profileRepo := profile.NewPostgresProfile(DB)

	articleUsecase := articleUsecase.NewArticleUsecase(articleRepo, sessionRepo)
	authUsecase := authUsecase.NewAuthUsecase(userRepo, sessionRepo)
	profileUsecase := profileUsecase.NewProfileUsecase(sessionRepo, userRepo, profileRepo)

	articleHandler := articleHandler.NewArticleHandler(articleUsecase)
	authHandler := authHandler.NewAuthHandler(authUsecase)
	profileHandler := profileHandler.NewProfileHandler(profileUsecase)

	mux := router.NewRouter(articleHandler, authHandler, profileHandler)
	handler := middleware.RecoverMiddleware(mux)

	server := http.Server{
		Addr:         config.BindAddr,
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	return &Server{
		config: config,
		server: server,
	}, nil
}

func (s *Server) StartServer() {
	logger.Info(nil, fmt.Sprintf("Starting MindLeak API server on %s", s.server.Addr))

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(nil, fmt.Sprintf("Server stopped with error: %v", err), nil)
	} else {
		logger.Info(nil, "Server stopped gracefully", nil)
	}
}
