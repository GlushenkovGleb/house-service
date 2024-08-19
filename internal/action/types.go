package action

import "github.com/GlushenkovGleb/house-service/internal/store"

type Store interface {
	store.House
	store.Flat
}
