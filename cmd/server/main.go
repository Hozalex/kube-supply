package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hozalex/kube-supply/internal/api"
)

func main() {

	s := &http.Server{
		Addr:              ":8080",
		Handler:           api.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	serverErrors := make(chan error, 1)
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		serverErrors <- s.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil {
			log.Fatal(err)
		}

	case <-ctx.Done():
		log.Println("Shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.Shutdown(shutdownCtx)
	if err != nil {
		if closeErr := s.Close(); closeErr != nil {
			log.Printf("Failed to close server: %v", closeErr)
		}
		os.Exit(1)
	}
}
