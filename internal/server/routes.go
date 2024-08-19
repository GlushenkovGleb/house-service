package server

import (
	"github.com/GlushenkovGleb/house-service/internal/handler"
	"github.com/go-chi/chi/v5"
)

func registerRoutes(r chi.Router, handler handler.Handler) {
	r.Post("/house/create", handler.CreateHouse)
	r.Post("/flat/create", handler.CreateFlat)
	r.Post("/flat/update", handler.UpdateFlat)
	r.Get("/house/{id}", handler.GetFlats)
}
