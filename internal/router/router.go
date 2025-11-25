package router

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/appeal"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/auth"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/categories"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/comment"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/subscriptions"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/swagger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/topBlogs"
	chat "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/ws"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(
	articleHandler *article.Handler,
	authHandler *auth.Handler,
	profileHandler *profile.Handler,
	categoryHandler *categories.Handler,
	topBlogsHandler *topBlogs.Handler,
	commentHandler *comment.Handler,
	subsHandler *subscriptions.Handler,
	appealHandler *appeal.Handler,
	chatHandler *chat.Handler,
) *mux.Router {
	router := mux.NewRouter()

	// Мидлвары
	router.Use(middleware.RecoverMiddleware)
	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.CORSMiddleware)
	// router.Use(middleware.CSRFMiddleware)

	// Авторизация
	router.HandleFunc("/feed", articleHandler.Feed).Methods("GET")
	router.HandleFunc("/registration", authHandler.Registration).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("GET")
	router.HandleFunc("/me", authHandler.Me).Methods("GET")

	// Сваггер
	router.HandleFunc("/swagger", swagger.SwaggerHandler).Methods("GET")

	// Профиль
	router.HandleFunc("/profile", profileHandler.ShowProfileHandler).Methods("GET")
	router.HandleFunc("/profile", profileHandler.EditProfileHandler).Methods("PUT")
	router.HandleFunc("/profile/delete", profileHandler.DeleteProfileHandler).Methods("DELETE")
	router.HandleFunc("/uploads/avatar", profileHandler.UploadAvatar).Methods("POST")
	router.HandleFunc("/delete/avatar", profileHandler.DeleteAvatar).Methods("DELETE")
	router.HandleFunc("/uploads/cover", profileHandler.UploadCover).Methods("POST")
	router.HandleFunc("/delete/cover", profileHandler.DeleteCover).Methods("DELETE")

	// Топ блогов
	router.HandleFunc("/topblogs", topBlogsHandler.ShowTopBlogs).Methods("GET")

	// Категории
	router.HandleFunc("/feed/category", categoryHandler.CategoriesHandler).Methods("GET")

	// Комменты
	router.HandleFunc("/comments", commentHandler.GetCommentsHandler).Methods("GET")
	router.HandleFunc("/comments", commentHandler.DeleteCommentHandler).Methods("DELETE")
	router.HandleFunc("/comments", commentHandler.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/comments", commentHandler.UpdateCommentHandler).Methods("PUT")

	// Посты
	router.HandleFunc("/posts", articleHandler.CreateArticle).Methods("POST")
	router.HandleFunc("/posts/{id}", articleHandler.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/posts/{id}", articleHandler.UpdateArticle).Methods("PUT")
	router.HandleFunc("/posts", articleHandler.GetArticlesByAuthorId).Methods("GET")
	router.HandleFunc("/post", articleHandler.GetArticle).Methods("GET")
	router.HandleFunc("/uploads/media", articleHandler.UploadMedia).Methods("POST")
	router.HandleFunc("/delete/media", articleHandler.DeleteMedia).Methods("DELETE")

	// Подписки
	router.HandleFunc("subscribers", subsHandler.GetSubscribers).Methods("GET")
	router.HandleFunc("subscriptions", subsHandler.GetSubscriptions).Methods("GET")
	router.HandleFunc("unsubscribe/{id}", subsHandler.Unsubscribe).Methods("POST")
	router.HandleFunc("subscribe/{id}", subsHandler.Subscribe).Methods("POST")

	//  Техподдержка (хакатон)
	router.HandleFunc("/appeal", appealHandler.CreateAppeal).Methods("POST")
	router.HandleFunc("/appeal/anonymous", appealHandler.CreateAnonymousAppeal).Methods("POST")
	router.HandleFunc("/appeal", appealHandler.DeleteAppeal).Methods("DELETE")
	router.HandleFunc("/appeal", appealHandler.GetAppealByID).Methods("GET")
	router.HandleFunc("/appeals/uploads", appealHandler.UploadScreenshot).Methods("POST")
	// router.HandleFunc("/delete/screenshot", appealHandler.DeleteMedia).Methods("DELETE")
	router.HandleFunc("/appeals", appealHandler.GetAppeals).Methods("GET")
	router.HandleFunc("/appeals/statistics", appealHandler.GetAppealsStatistics).Methods("GET")
	router.HandleFunc("appeals/categories", appealHandler.GetCategory).Methods("GET")

	// чаты
	router.HandleFunc("/chat/rooms", chatHandler.GetRooms).Methods("GET")
	router.HandleFunc("/chat/rooms", chatHandler.CreateRoom).Methods("POST")
	router.HandleFunc("/chat/rooms/{room_id}/messages", chatHandler.GetMessages).Methods("GET")
	router.HandleFunc("/chat/ws", chatHandler.Connect).Methods("GET")
	return router
}
