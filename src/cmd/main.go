package main

import (
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/server"
	serverConfig "github/ggualbertosouza/Karhub-Desafio-Backend/src/server/config"
	"log"
)

func main() {
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

	router := server.NewRouter(cfg.App.Environment)
	srv := server.NewServer(router, cfg)

	srv.Start()
}
