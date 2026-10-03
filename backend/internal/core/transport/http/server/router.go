package core_transport_http_server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"
)

func NewRouter(pool *core_postgres_pool.ConnectionPool, log *slog.Logger) chi.Router {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "multipart/form-data"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	Routes(pool, router, log)

	ImageFs := http.FileServer(http.Dir("./public/images"))

	router.Handle("/images/*", http.StripPrefix("/images/", ImageFs))

	PdfFs := http.FileServer(http.Dir("./public/documents"))

	router.Handle("/documents/*", http.StripPrefix("/documents/", PdfFs))

	return router
}
