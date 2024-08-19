package dto

import (
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"net/http"
	"time"
)

type CreateHouseResponse struct {
	ID        int       `json:"id"`
	Address   string    `json:"address"`
	Year      int       `json:"year"`
	Developer string    `json:"developer,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *CreateHouseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func ToCreateHouseResponse(h *domain.House) *CreateHouseResponse {
	return &CreateHouseResponse{
		ID:        int(h.ID.Int64),
		Address:   h.Address,
		Year:      h.Year,
		Developer: h.Developer.String,
		CreatedAt: h.CreatedAt.Time,
		UpdatedAt: h.UpdatedAt.Time,
	}
}
