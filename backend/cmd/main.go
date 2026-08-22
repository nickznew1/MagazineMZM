package main

import (
	_ "github.com/lib/pq"
	"github.com/nickznew1/MagazineMZM/backend/internal/config"
	"github.com/nickznew1/MagazineMZM/backend/internal/routes"
	"github.com/nickznew1/MagazineMZM/backend/storage"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad("config/config-example.yaml")

	logger := setupLogger(cfg.Env)

	db, err := storage.Connect(cfg)
	if err != nil {
		logger.Error("Error when connecting to database psql", slog.Any("error", err))
		os.Exit(1)
	}

	logger.Info("Starting backend MZM app", slog.String("env", cfg.Env))
	logger.Debug("Debug messages are enabled")

	routes.Routes(db, cfg, logger)

}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "local":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
