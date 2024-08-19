package server

import (
	"github.com/GlushenkovGleb/house-service/internal/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func New(handler handler.Handler) http.Server {
	r := chi.NewRouter()
	registerRoutes(r, handler)

	return http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
