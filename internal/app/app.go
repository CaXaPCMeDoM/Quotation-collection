package app

import (
	"citatnik/config"
	"citatnik/internal/controller/http"
	"citatnik/internal/pkg/httpserver"
	"citatnik/internal/repo/memory"
	"citatnik/internal/usecase/quote"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	//Repository
	quoteRepo := memory.New()

	// UserCase
	quoteUC := quote.New(quoteRepo)

	// HTTP
	router := http.NewRouter(quoteUC)

	// Server
	server := httpserver.New(
		cfg,
		router,
	)

	server.Start()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - Run - signal: %s", s.String())
	case err := <-server.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := server.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
