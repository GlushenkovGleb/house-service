package dto

import (
	"errors"
	"fmt"
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"net/http"
)

type CreateHouseRequest struct {
	Address   string `json:"address"`
	Year      int    `json:"year"`
	Developer string `json:"developer"`
}

func (c *CreateHouseRequest) IsValid() bool {
	if c.Address == "" || c.Year == 0 {
		return false
	}

	return true
}

func (c *CreateHouseRequest) ToHouse() *domain.House {
	return &domain.House{
		Address:   c.Address,
		Year:      c.Year,
		Developer: domain.AutoString(c.Developer),
	}
}

func (c *CreateHouseRequest) Bind(r *http.Request) error {
	fmt.Println(fmt.Sprintf("I'm checking this request body %v\n", c))
	if !c.IsValid() {
		return errors.New("invalid request")
	}
	return nil
}
