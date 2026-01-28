package serverConfig

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() (*EnvConfig, error) {
	initEnv()

	cfg := &EnvConfig{}
	setDefaults(cfg)

	if err := validate(cfg); err != nil {
		return nil, err
	}

	log.Printf("Environment variables setted up, using %s Env", cfg.App.Environment)
	return cfg, nil
}

func initEnv() {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development")
	}

	_ = godotenv.Load(".env")
}

func setDefaults(cfg *EnvConfig) {
	cfg.App.Environment = os.Getenv("APP_ENV")

	cfg.App.Host = os.Getenv("APP_HOST")
	if cfg.App.Host == "" {
		cfg.App.Host = "0.0.0.0"
	}

	portStr := os.Getenv("APP_PORT")
	if portStr != "" {
		port, _ := strconv.Atoi(portStr)
		cfg.App.Port = port
	}
}

func validate(cfg *EnvConfig) error {
	if cfg.App.Port == 0 {
		return errors.New("APP_PORT is required")
	}

	if cfg.App.Environment == "" {
		return errors.New("APP_ENV is required")
	}

	return nil
}
