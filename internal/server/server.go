package server

import (
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/middleware"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/router"
)

func StartServer() {
	sessions := session.NewInMemorySession()
	users := repository.NewInMemoryUser()
	articles := repository.NewInMemoryArticle()

	mux := router.NewRouter(sessions, users, articles)
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
