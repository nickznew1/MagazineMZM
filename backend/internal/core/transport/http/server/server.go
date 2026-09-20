package core_transport_http_server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type HTTPServer struct {
	config Config
	log    *slog.Logger
}

func NewHTTPServer(
	config Config,
	log *slog.Logger,
) *HTTPServer {

	return &HTTPServer{
		config: config,
		log:    log,
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := http.Server{
		Addr: h.config.Port,
	}

	ch := make(chan error, 1)

	go func() {
		// LOGGER
		err := server.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				ch <- err
			}
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Listen and serve HTTP-server: %w", err)
		}
	case <-ctx.Done():
		//LOGGER

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("Shutdown HTTP server: %w", err)
		}
		//LOGGER
	}
	return nil
}
