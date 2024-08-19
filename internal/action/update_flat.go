package action

import "github.com/GlushenkovGleb/house-service/internal/domain"

func (c *Client) UpdateFlat(id int, status string) (*domain.Flat, error) {
	return c.store.UpdateFlat(id, status)
}
