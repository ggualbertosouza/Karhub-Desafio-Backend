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
	// Application
	cfg.App.Environment = os.Getenv("APP_ENV")

	cfg.App.Host = os.Getenv("APP_HOST")
	if cfg.App.Host == "" {
		cfg.App.Host = "0.0.0.0"
	}

	if portStr := os.Getenv("APP_PORT"); portStr != "" {
		port, _ := strconv.Atoi(portStr)
		cfg.App.Port = port
	}

	// DB
	cfg.Db.Host = os.Getenv("DB_HOST")
	if cfg.Db.Host == "" {
		cfg.Db.Host = "localhost"
	}

	if portStr := os.Getenv("DB_PORT"); portStr != "" {
		port, _ := strconv.Atoi(portStr)
		cfg.Db.Port = port
	} else {
		cfg.Db.Port = 5432
	}

	cfg.Db.Name = os.Getenv("DB_NAME")
	cfg.Db.User = os.Getenv("DB_USER")
	cfg.Db.Password = os.Getenv("DB_PASSWORD")

	cfg.Db.SSLMode = os.Getenv("DB_SSLMODE")
	if cfg.Db.SSLMode == "" {
		cfg.Db.SSLMode = "disable"
	}

	cfg.Mocks.Path = os.Getenv("PLAYLIST_MOCK_PATH")
}

func validate(cfg *EnvConfig) error {
	// ---- APP ----
	if cfg.App.Port == 0 {
		return errors.New("APP_PORT is required")
	}

	if cfg.App.Environment == "" {
		return errors.New("APP_ENV is required")
	}

	// ---- DB ----
	if cfg.Db.Host == "" {
		return errors.New("DB_HOST is required")
	}

	if cfg.Db.Port == 0 {
		return errors.New("DB_PORT is required")
	}

	if cfg.Db.Name == "" {
		return errors.New("DB_NAME is required")
	}

	if cfg.Db.User == "" {
		return errors.New("DB_USER is required")
	}

	if cfg.Db.Password == "" {
		return errors.New("DB_PASSWORD is required")
	}

	return nil
}
