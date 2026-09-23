package core_logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func SetupLogger(config Config) *slog.Logger {
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

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value("log").(*slog.Logger)

	if !ok {
		panic("No logger in context")
	}

	return logger
}
