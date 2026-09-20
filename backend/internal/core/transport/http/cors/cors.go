package core_transport_http_cors

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewCORS(config Config, router chi.Router) {

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.Addr},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "multipart/form-data"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	
}
