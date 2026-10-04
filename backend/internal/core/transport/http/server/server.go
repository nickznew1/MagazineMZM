package core_transport_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
)

type HTTPServer struct {
	config Config
	log    *core_logger.Logger
	router chi.Router
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
) *HTTPServer {

	return &HTTPServer{
		config: config,
		log:    log,
	}
}

func (h *HTTPServer) Run(ctx context.Context, router chi.Router) error {
	server := http.Server{
		Addr:    h.config.Port,
		Handler: router,
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
