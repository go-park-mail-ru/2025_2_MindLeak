package apiserver

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/minio"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/postgres"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/redis"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	articleUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	authUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"

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
	server http.Server
}

func New(config *server.Config) (*Server, error) {
	logger.Info(nil, "Logger initialized")
	logger.Info(nil, "Server started initializing")

	//INITIALIZE POSTGRES
	PGConfig := postgres.NewPostgresConfig()
	DB, err := PGConfig.PGconnect()
	if err != nil {
		logger.Error(nil, "Error initializing DB connection")
		return nil, err
	}

	logger.Info(nil, "Postgres connection initialized")

	//INITIALIZE REDIS
	RedisConfig := redis.NewRedisConfig()
	RedisConn, err := RedisConfig.RedisConnect()
	if err != nil {
		logger.Error(nil, "Error initializing Redis connection")
		//return nil, err
	}

	//INITIALIZE MINIO
	_ = minio.NewMinioConfig()

	logger.Info(nil, "Redis connection initialized")

	sessionRepo := session.NewRedisSessionManager(RedisConn)
	userRepo := user.NewPostgresUser(DB)
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
		server: server,
	}, nil
}

func (s *Server) StartServer() {
	logger.Info(nil, "Starting MindLeak API server")
	s.server.ListenAndServe()
}
