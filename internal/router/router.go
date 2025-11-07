package router

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/auth"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/categories"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/comment"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/swagger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/topBlogs"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(
	articleHandler *article.Handler,
	authHandler *auth.Handler,
	profileHandler *profile.Handler,
	categoryHandler *categories.Handler,
	subsHandler *topBlogs.Handler,
	commentHandler *comment.Handler,
) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.RecoverMiddleware)
	router.Use(middleware.RequestIDMiddleware)
	router.Use(middleware.CORSMiddleware)
	//router.Use(middleware.CSRFMiddleware)

	//router.Use(middleware.AuthMiddleware) Потом подключить к нужным ручкам

	router.HandleFunc("/feed", articleHandler.Feed).Methods("GET")
	router.HandleFunc("/registration", authHandler.Registration).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("GET")
	router.HandleFunc("/me", authHandler.Me).Methods("GET")
	router.HandleFunc("/swagger", swagger.SwaggerHandler).Methods("GET")

	router.HandleFunc("/profile/{id}", profileHandler.ShowOtherProfileHandler).Methods("GET") //чужой профиль
	router.HandleFunc("/profile", profileHandler.ShowOwnProfileHandler).Methods("GET")        //свой профиль
	router.HandleFunc("/profile", profileHandler.EditProfileHandler).Methods("PUT")
	router.HandleFunc("/profile/delete", profileHandler.DeleteProfileHandler).Methods("DELETE")

	router.HandleFunc("/uploads/avatar", profileHandler.UploadAvatar).Methods("POST")
	router.HandleFunc("/delete/avatar", profileHandler.DeleteAvatar).Methods("DELETE")
	router.HandleFunc("/uploads/cover", profileHandler.UploadCover).Methods("POST")
	router.HandleFunc("/delete/cover", profileHandler.DeleteCover).Methods("DELETE")

	router.HandleFunc("/topblogs", subsHandler.ShowTopBlogs).Methods("GET")

	router.HandleFunc("/feed/category", categoryHandler.CategoriesHandler).Methods("GET")

	// Комменты
	router.HandleFunc("/comments", commentHandler.GetCommentsHandler).Methods("GET")
	router.HandleFunc("/comments", commentHandler.DeleteCommentHandler).Methods("DELETE")
	router.HandleFunc("/comments", commentHandler.CreateCommentHandler).Methods("POST")
	router.HandleFunc("/comments", commentHandler.UpdateCommentHandler).Methods("PUT")

	//Посты
	router.HandleFunc("/posts", articleHandler.CreateArticle).Methods("POST")
	router.HandleFunc("/posts/{id}", articleHandler.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/posts/{id}", articleHandler.UpdateArticle).Methods("PUT")
	router.HandleFunc("/posts", articleHandler.GetArticlesByAuthorId).Methods("GET")
	router.HandleFunc("/post", articleHandler.GetArticle).Methods("GET")
	router.HandleFunc("/uploads/media", articleHandler.UploadMedia).Methods("POST")
	router.HandleFunc("/delete/media", articleHandler.DeleteMedia).Methods("DELETE")

	return router
}
