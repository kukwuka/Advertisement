package teststore

import (
	"Advertising/internal/app/model"
	"Advertising/internal/app/store"
)

//Store ...
type Store struct {
	AdvertisementRepository *AdvertisementRepository
}

//New ...
func New() *Store {
	return &Store{}
}

//Statistic ...
func (s *Store) Advertisement() store.AdvertisementRepository {
	if s.AdvertisementRepository != nil {
		return s.AdvertisementRepository
	}
	s.AdvertisementRepository = &AdvertisementRepository{
		store: s,
		users: make(map[int]*model.Advertisement),
	}
	return s.AdvertisementRepository
}
