package main

import (
	"log"
	"net/http"
	"time"

	"github.com/runAlgo/go-auth/internal/httpserver"
)

func main() {

	router := httpserver.NewRouter()

	// standard go type that runs a http server
	srv := &http.Server{
		Addr:              ":5000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Api running on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("server closed")
			return
		}
		log.Fatalf("server error: %v", err)
	}
}
