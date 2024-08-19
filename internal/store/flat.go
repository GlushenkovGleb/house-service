package store

import (
	"github.com/GlushenkovGleb/house-service/internal/domain"
	"github.com/Masterminds/squirrel"
)

type Flat interface {
	CreateFlat(flat *domain.Flat) error
	UpdateFlat(id int, status string) (*domain.Flat, error)
	GetFlats(houseID int) ([]domain.Flat, error)
}

var _ Flat = (*Store)(nil)

func (s *Store) CreateFlat(flat *domain.Flat) error {
	q, args, err := Insert("flats", flat).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return err
	}

	return s.db.Get(flat, q, args...)
}

func (s *Store) UpdateFlat(id int, status string) (*domain.Flat, error) {
	q, args, err := psql.Update("flats").
		Where(squirrel.Eq{"id": id}).
		Set("status", status).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	res := &domain.Flat{}
	err = s.db.Get(res, q, args...)
	return res, err
}

func (s *Store) GetFlats(houseID int) ([]domain.Flat, error) {
	//TODO implement me
	panic("implement me")
}
