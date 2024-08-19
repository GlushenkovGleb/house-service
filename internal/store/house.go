package store

import (
	"github.com/GlushenkovGleb/house-service/internal/domain"
)

type House interface {
	CreateHouse(house *domain.House) error
}

var _ House = (*Store)(nil)

func (s *Store) CreateHouse(house *domain.House) error {
	q, args, err := Insert("houses", house).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return err
	}

	return s.db.Get(house, q, args...)
}
