package action

import (
	"github.com/GlushenkovGleb/house-service/internal/domain"
)

func (c *Client) CreateHouse(house *domain.House) error {
	return c.store.CreateHouse(house)
}
