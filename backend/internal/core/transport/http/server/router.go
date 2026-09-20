package core_transport_http_server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	core_transport_http_cors "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/cors"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	configCORS := core_transport_http_cors.NewConfigMust()

	core_transport_http_cors.NewCORS(configCORS, router)

	ImageFs := http.FileServer(http.Dir("./public/images"))

	router.Handle("/images/*", http.StripPrefix("/images/", ImageFs))

	PdfFs := http.FileServer(http.Dir("./public/documents"))

	router.Handle("/documents/*", http.StripPrefix("/documents/", PdfFs))

	return router
}
