package server

import (
	"context"
	"fmt"
	serverConfig "github/ggualbertosouza/Karhub-Desafio-Backend/src/server/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	httpServer *http.Server
	config     *serverConfig.EnvConfig
}

func NewServer(router *gin.Engine, cfg *serverConfig.EnvConfig) *Server {
	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)

	return &Server{
		router: router,
		config: cfg,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) Start() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("Server running")

		err := s.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting server: %v", err)
		}
	}()

	<-stop
	log.Println("Shutdown signal received...")
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Graceful shutdown error: %v", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}
