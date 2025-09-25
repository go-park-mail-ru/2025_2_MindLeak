package server

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL:", r.URL.String())
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
