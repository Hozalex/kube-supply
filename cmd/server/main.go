package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Hozalex/kube-supply/internal/api"
)

func main() {
	s := &http.Server{
		Addr:              ":8080",
		Handler:           api.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
