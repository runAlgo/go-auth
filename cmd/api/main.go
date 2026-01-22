package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/runAlgo/go-auth/internal/app"
	"github.com/runAlgo/go-auth/internal/httpserver"
)

func main() {
	// root context
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Startup failed: %v", err)
	}

	defer func(){
		if err := a.Close(ctx); err != nil {
			log.Printf("Shutdown warning: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)

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
