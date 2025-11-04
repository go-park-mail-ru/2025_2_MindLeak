package apiserver

import (
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/minio"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/postgres"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/redis"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	articleUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	authUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"
	categoryUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/categories/usecase"
	profileUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/usecase"

	articleHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	authHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	categoryHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/categories"
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

	if err := RedisConfig.Ping(); err != nil {
		logger.Error(nil, "Error initializing Redis connection: %v", err)
		return nil, err
	}

	logger.Info(nil, "Redis pool initialized")

	//INITIALIZE MINIO
	minioCfg := minio.NewMinioConfig()
	minioClient, err := minio_client.NewClient(
		minioCfg.Endpoint,
		minioCfg.User,
		minioCfg.Password,
		minioCfg.Bucket,
		minioCfg.Bucket)
	if err != nil {
		logger.Error(nil, "Error initializing MinIO client")
		return nil, err
	}
	logger.Info(nil, "MinIO client initialized")

	sessionRepo := session.NewRedisSessionManager(RedisConfig.GetPool())
	userRepo := user.NewPostgresUser(DB, minioClient)
	articleRepo := article.NewArticleRepo(DB)
	profileRepo := profile.NewPostgresProfile(DB, minioClient)

	articleUsecase := articleUsecase.NewArticleUsecase(articleRepo, sessionRepo)
	profileUsecase := profileUsecase.NewProfileUsecase(sessionRepo, userRepo, profileRepo, minioClient)
	authUsecase := authUsecase.NewAuthUsecase(userRepo, sessionRepo, profileRepo)
	categoryUsecase := categoryUsecase.NewCategoriesUsecase(articleRepo)

	articleHandler := articleHandler.NewArticleHandler(articleUsecase)
	authHandler := authHandler.NewAuthHandler(authUsecase)
	profileHandler := profileHandler.NewProfileHandler(profileUsecase)
	categoryHandler := categoryHandler.NewCategoriesHandler(categoryUsecase)

	mux := router.NewRouter(articleHandler, authHandler, profileHandler, categoryHandler)

	handler := middleware.CORSMiddleware(mux)
	handler = middleware.RecoverMiddleware(handler)

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
		logger.Error(nil, "Server stopped with error: %v", err)
	} else {
		logger.Info(nil, "Server stopped gracefully")
	}
}
