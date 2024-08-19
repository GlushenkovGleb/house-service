package dto

import (
	"errors"
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"net/http"
)

type CreateFlatRequest struct {
	HouseID int `json:"house_id"`
	Number  int `json:"number"`
	Price   int `json:"price"`
	Rooms   int `json:"rooms"`
}

func (cr *CreateFlatRequest) IsValid() bool {
	if cr.HouseID <= 0 || cr.Price <= 0 || cr.Rooms <= 0 || cr.Number <= 0 {
		return false
	}

	return true
}

func (cr *CreateFlatRequest) Bind(r *http.Request) error {
	if !cr.IsValid() {
		return errors.New("invalid request")
	}

	return nil
}

func (cr *CreateFlatRequest) ToFlat() *domain.Flat {
	return &domain.Flat{
		HouseID: cr.HouseID,
		Price:   cr.Price,
		Rooms:   cr.Rooms,
	}
}
