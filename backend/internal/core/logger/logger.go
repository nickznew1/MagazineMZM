package core_logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	*slog.Logger
	file *os.File
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(args...),
		file:   l.file,
	}
}

func SetupLogger(config Config) (*Logger, error) {

	var logger *slog.Logger

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil,
			fmt.Errorf("failed to create out folder for logger: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.00000")

	logFilePath := filepath.Join(
		config.Folder,
		fmt.Sprintf("%s.log", timestamp),
	)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil,
			fmt.Errorf("failed to open log file: %w", err)
	}

	var loggerLevel slog.Level

	switch config.Level {
	case "DEBUG":
		loggerLevel = slog.LevelDebug
	case "dev":
		loggerLevel = slog.LevelDebug
	case "prod":
		loggerLevel = slog.LevelInfo
	default:
		logFile.Close()
		return nil, fmt.Errorf("unknown loglevel: %w", err)
	}

	logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.MessageKey {
				a.Key = "MESSAGE"
			}
			if a.Key == slog.SourceKey {
				a.Value = slog.StringValue("SOURCE")
			}
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05"))
			}
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				switch level {
				case slog.LevelDebug:
					a.Value = slog.StringValue("DBG")
				case slog.LevelInfo:
					a.Value = slog.StringValue("INF")
				case slog.LevelWarn:
					a.Value = slog.StringValue("WRN")
				case slog.LevelError:
					a.Value = slog.StringValue("ERR")
				}
			}
			return a
		},
		Level: loggerLevel,
	}))

	return &Logger{
		logger,
		logFile,
	}, nil
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value("log").(*Logger)

	if !ok {
		panic("No logger in context")
	}

	return logger
}

func (l *Logger) Close() error {
	if err := l.file.Close(); err != nil {
		return fmt.Errorf("failed to close log file: %w")
	}
	return nil
}
