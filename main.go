package main

import (
	"context"
	"goevent/database"
	"goevent/domain"
	"goevent/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	newRouter := router.NewRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: newRouter,
	}

	domain.InitBusses()

	err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	log.Print("\nServer started on port " + srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
