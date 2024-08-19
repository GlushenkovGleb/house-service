package action

import "github.com/GlushenkovGleb/house-service/internal/domain"

func (c *Client) CreateFlat(flat *domain.Flat) error {
	return c.store.CreateFlat(flat)
}
