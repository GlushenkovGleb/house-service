package dto

import (
	"errors"
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"net/http"
)

type UpdateFlatRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (cr *UpdateFlatRequest) IsValid() bool {
	if cr.ID <= 0 || !domain.IsValidFlatStatus(cr.Status) {
		return false
	}

	return true
}

func (cr *UpdateFlatRequest) Bind(r *http.Request) error {
	if !cr.IsValid() {
		return errors.New("invalid request")
	}

	return nil
}
