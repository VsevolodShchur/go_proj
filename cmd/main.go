package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"proj/internal/repository"
	"proj/internal/service"
	"proj/internal/transport/http"
	"proj/pkg/server"
	"syscall"
)

func main() {
	repos := repository.Init()
	services := service.Init(repos)
	handlers := http.NewHandlers(services)
	r := http.NewRouter(handlers)

	s := server.NewServer("8080", r.Init())

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
