package main

import (
	"context"
	InMemoryCache "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/cache/inMemory"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/server"
	serverConfig "github/ggualbertosouza/Karhub-Desafio-Backend/src/server/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Environment config
	cfg, err := serverConfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error while starting environment variables: %v", err)
	}

	// Db config
	err = postgres.Connect(postgres.Config{
		Host:     cfg.Db.Host,
		Port:     cfg.Db.Port,
		User:     cfg.Db.User,
		Password: cfg.Db.Password,
		DbName:   cfg.Db.Name,
		SSLMode:  cfg.Db.SSLMode,
	})
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	// Cache init
	if err := InMemoryCache.InitBeerStyleCache(ctx); err != nil {
		log.Fatal(err)
	}

	if err := InMemoryCache.InitPlaylistMockCache(cfg.Mocks.Path); err != nil {
		log.Fatal(err)
	}

	router := server.NewRouter(cfg.App.Environment)
	srv := server.NewServer(router, cfg)

	go func() {
		if err := srv.Start(ctx); err != nil {
			log.Println(err)
		}
	}()

	<-stop
	log.Println("Shutdown signal received")

	cancel()
	InMemoryCache.CloseBeerStyleCache()
}
