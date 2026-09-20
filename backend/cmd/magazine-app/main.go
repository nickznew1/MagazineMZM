package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/lmittmann/tint"
	"github.com/nickznew1/MagazineMZM/backend/internal/config"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"
	core_transport_http_server "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/server"
	"github.com/nickznew1/MagazineMZM/backend/internal/routes"
	"github.com/nickznew1/MagazineMZM/backend/storage"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)

	defer cancel()

	logger := setupLogger(
		core_logger.NewConfigMust(),
	)

	pool, err := core_postgres_pool.NewConnectionPool(
		core_postgres_pool.NewConfigMust(),
		ctx)

	if err != nil {
		logger.Error("Failed to init postgres pool", error.Error(err))
	}

	defer pool.Close()

	logger.Info("Starting backend MZM app")
	logger.Debug("Debug messages are enabled")

	httpServer := core_transport_http_server.NewHTTPServer(
		core_transport_http_server.NewConfigMust(),
		logger,
	)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("Failed to run http server", error.Error(err))
	}

	router := core_transport_http_server.NewRouter()

	routes.Routes(db, cfg, logger)

}

func setupLogger(config core_logger.Config) *slog.Logger {
	var logger *slog.Logger

	switch config.Level {
	case "local":
		logger = slog.New(tint.NewTextHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))
	case "dev":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
