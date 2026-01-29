package server

import (
	"context"
	"fmt"
	serverConfig "github/ggualbertosouza/Karhub-Desafio-Backend/src/server/config"
	"log"
	"net/http"
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

func (s *Server) Start(ctx context.Context) error {
	go func() {
		log.Printf("Server running: %s:%d", s.config.App.Host, s.config.App.Port)

		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server context cancelled")

	return s.Shutdown()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Graceful shutdown error: %v", err)
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
