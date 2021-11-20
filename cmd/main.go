package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"proj/internal/transport/http"
	"proj/pkg/server"
	"syscall"
)

func main() {
	r := http.NewRouter()
	s := server.NewServer("8080", r)

	go func() {
		if err := s.Start(); err != nil {
			log.Printf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}
