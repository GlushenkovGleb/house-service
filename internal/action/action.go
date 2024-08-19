package action

import (
	"github.com/GlushenkovGleb/house-service/internal/domain"
)

type Action interface {
	CreateHouse(house *domain.House) error
	CreateFlat(flat *domain.Flat) error
	UpdateFlat(id int, status string) (*domain.Flat, error)
	// GetFlats TODO: instead of houseID use get_flats_opts
	GetFlats(houseID int) ([]domain.Flat, error)
}

type Client struct {
	store Store
}

func (c *Client) GetFlats(houseID int) ([]domain.Flat, error) {
	//TODO implement me
	panic("implement me")
}

var _ Action = (*Client)(nil)

func New(store Store) *Client {
	return &Client{
		store: store,
	}
}
