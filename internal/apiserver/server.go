package apiserver

import (
	"fmt"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/minio"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/postgres"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/config/redis"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/appeal"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/comment"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/subscriptions"

	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	appealUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/appeal/usecase"
	articleUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/usecase"
	authUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/usecase"
	categoryUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/categories/usecase"
	commentUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/usecase"
	profileUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/usecase"
	subsUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/subscriptions/usecase"
	topBlogsUsecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/topBlogs/usecase"

	appealHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal"
	articleHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	authHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	categoryHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/categories"
	commentHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/comment"
	profileHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile"
	subsHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/subscriptions"
	topBlogsHandler "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/topBlogs"

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
	subsRepo := subscriptions.NewPostgresSubscription(DB)
	commentRepo := comment.NewPostgresComment(DB)
	appealRepo := appeal.NewPostgresAppeal(DB)

	articleUsecase := articleUsecase.NewArticleUsecase(articleRepo, sessionRepo, minioClient)
	profileUsecase := profileUsecase.NewProfileUsecase(sessionRepo, userRepo, profileRepo, minioClient)
	authUsecase := authUsecase.NewAuthUsecase(userRepo, sessionRepo, profileRepo)
	categoryUsecase := categoryUsecase.NewCategoriesUsecase(articleRepo)
	topBlogsUsecase := topBlogsUsecase.NewTopBlogsUsecase(subsRepo)
	commentUsecase := commentUsecase.NewCommentUsecase(commentRepo, sessionRepo)
	subsUsecase := subsUsecase.NewSubscriptionsUsecase(subsRepo, sessionRepo, userRepo)
	appealUsecase := appealUsecase.NewAppealUsecase(appealRepo, sessionRepo, minioClient)

	articleHandler := articleHandler.NewArticleHandler(articleUsecase)
	authHandler := authHandler.NewAuthHandler(authUsecase)
	profileHandler := profileHandler.NewProfileHandler(profileUsecase)
	categoryHandler := categoryHandler.NewCategoriesHandler(categoryUsecase)
	topBlogsHandler := topBlogsHandler.NewTopBlogsHandler(topBlogsUsecase)
	commentHandler := commentHandler.NewCommentHandler(commentUsecase)
	subsHandler := subsHandler.NewSubsHandler(subsUsecase)
	appealHandler := appealHandler.NewAppealHandler(appealUsecase)

	mux := router.NewRouter(
		articleHandler,
		authHandler,
		profileHandler,
		categoryHandler,
		topBlogsHandler,
		commentHandler,
		subsHandler,
		appealHandler,
	)

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
