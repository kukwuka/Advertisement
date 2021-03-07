package teststore_test

import (
	"Advertising/internal/app/model"
	"Advertising/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvertisementRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestAdvertisement(t)
	Id, err := s.Advertisement().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, Id)
}

func TestAdvertisementRepository_GetOne(t *testing.T) {
	s := teststore.New()
	u1 := model.TestAdvertisement(t)
	Id, _ := s.Advertisement().Create(u1)
	u2, err := s.Advertisement().GetOne(Id)
	assert.NoError(t, err)
	assert.NotNil(t, u2)

}

func TestStatisticRepository_GetAll(t *testing.T) {
	s := teststore.New()
	u1 := model.TestAdvertisement(t)
	s.Advertisement().Create(u1)
	u2, err := s.Advertisement().GetAll(0)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestStatisticRepository_GetAllMap(t *testing.T) {
	s := teststore.New()
	u1 := model.TestAdvertisement(t)
	s.Advertisement().Create(u1)
	u2, err := s.Advertisement().GetAllMap(0)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
