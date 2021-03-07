package sqlstore

import (
	"Advertising/internal/app/store"
	"database/sql"

	_ "github.com/lib/pq" // ...
)

//Store ...
type Store struct {
	db                      *sql.DB
	advertisementRepository *AdvertisementRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Statistic ...
func (s *Store) Advertisement() store.AdvertisementRepository {
	if s.advertisementRepository != nil {
		return s.advertisementRepository
	}
	s.advertisementRepository = &AdvertisementRepository{
		store: s,
	}
	return s.advertisementRepository
}
