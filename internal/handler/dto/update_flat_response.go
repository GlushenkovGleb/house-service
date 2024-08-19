package dto

import (
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"net/http"
)

type UpdateFlatResponse struct {
	ID      int               `json:"id"`
	Number  int               `json:"number"`
	HouseID int               `json:"house_id"`
	Price   int               `json:"price"`
	Rooms   int               `json:"rooms"`
	Status  domain.FlatStatus `json:"status"`
}

func (cr *UpdateFlatResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	return nil
}

func ToUpdateFlatResponse(flat *domain.Flat) *UpdateFlatResponse {
	return &UpdateFlatResponse{
		ID:      int(flat.ID.Int64),
		Number:  flat.Number,
		HouseID: flat.HouseID,
		Price:   flat.Price,
		Rooms:   flat.Rooms,
		Status:  flat.Status,
	}
}
