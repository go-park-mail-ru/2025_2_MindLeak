package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/router"
)

func StartServer() {
	sessions := repository.NewInMemorySession()
	users := repository.NewInMemoryUser()
	articles := repository.NewInMemoryArticle()

	mux := router.NewRouter(sessions, users, articles)

	server := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8090")
	server.ListenAndServe()
}
