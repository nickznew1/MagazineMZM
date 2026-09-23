package main

import (
	"context"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"
	core_transport_http_server "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/server"
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

	logger := core_logger.SetupLogger(
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

	core_transport_http_server.Routes(pool, router, logger)
}
