package sqlstore_test

import (
	"Advertising/internal/app/model"
	"Advertising/internal/app/store/sqlstore"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAdvertisementRepository_Create(t *testing.T) {
	db, _ := sqlstore.TestDb(t, databaseURL)
	//defer teardown("advertisement")
	s := sqlstore.New(db)
	_, err := s.Advertisement().Create(&model.Advertisement{
		Name:        "первое",
		Description: "описание",
		Cost:        40,
		ImgUrl: pq.StringArray{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
			"https://i.imgur.com/ExdKOOz.png",
			"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
	})
	assert.NoError(t, err)
}

//
func TestAdvertisementRepository_GetAll(t *testing.T) {
	db, _ := sqlstore.TestDb(t, databaseURL)
	//defer teardown("advertisement")
	s := sqlstore.New(db)

	_, err := s.Advertisement().GetAll(1)

	assert.NoError(t, err)
}
func TestAdvertisementRepository_GetAllMap(t *testing.T) {
	db, _ := sqlstore.TestDb(t, databaseURL)
	//defer teardown("advertisement")
	s := sqlstore.New(db)

	m, err := s.Advertisement().GetAllMap(1)
	log.Print(m)
	assert.NoError(t, err)
}

func TestAdvertisementRepository_GetOne(t *testing.T) {
	db, _ := sqlstore.TestDb(t, databaseURL)
	//defer teardown("advertisement")
	s := sqlstore.New(db)
	m, err := s.Advertisement().GetAll(0)
	if err != nil {
		assert.Error(t, err)
	}
	_, err = s.Advertisement().GetOne(m[0].Id)
	assert.NoError(t, err)
}
