package main

import (
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/server"
	serverConfig "github/ggualbertosouza/Karhub-Desafio-Backend/src/server/config"
	"log"
)

func main() {
	cfg, err := serverConfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error while starting environment variables: %v", err)
	}

	router := server.NewRouter(cfg.App.Environment)
	srv := server.NewServer(router, cfg)

	srv.Start()
}
