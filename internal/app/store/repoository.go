package store

import (
	"Advertising/internal/app/model"
)

//StatisticRepository ...
type AdvertisementRepository interface {
	Create(*model.Advertisement) (int, error)
	GetAll(Page int) ([]*model.Advertisement, error)
	GetOne(Id int) (*model.Advertisement, error)
	GetAllMap(Page int) ([]map[string]interface{}, error)
}
