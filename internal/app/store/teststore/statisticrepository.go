package teststore

import (
	"Advertising/internal/app/model"
	"errors"
)

// UserRepository ...
type AdvertisementRepository struct {
	store *Store
	users map[int]*model.Advertisement
}

// Create ...
func (r *AdvertisementRepository) Create(u *model.Advertisement) (Id int, err error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}
	u.Id = len(r.users) + 1
	r.users[u.Id] = u

	return u.Id, nil
}

// GetOne ...
func (r *AdvertisementRepository) GetOne(Id int) (*model.Advertisement, error) {
	u, ok := r.users[Id]
	if !ok {
		return nil, errors.New("record not found")
	}
	return u, nil
}

//GetAllMap
func (r *AdvertisementRepository) GetAllMap(Page int) ([]map[string]interface{}, error) {
	us := make([]map[string]interface{}, 0, 10)
	for i := 0; i < Page; i++ {
		model := r.users[(Page-1)*10+i]
		u := map[string]interface{}{
			"id":   model.Id,
			"name": model.Name,
			"cost": model.Cost,
			"img":  model.ImgUrl[0],
		}
		us = append(us, u)
	}
	return us, nil
}
func (r *AdvertisementRepository) GetAll(Page int) ([]*model.Advertisement, error) {
	us := make([]*model.Advertisement, 0, 10)
	for i := 0; i < Page; i++ {
		model := r.users[(Page-1)*10+i]
		us = append(us, model)
	}
	return us, nil
}
