package model

import (
	"github.com/lib/pq"
	"testing"
)

// TestUser ...
func TestAdvertisement(t *testing.T) *Advertisement {
	t.Helper()

	return &Advertisement{
		Name:        "первое",
		Description: "описание",
		Cost:        40,
		ImgUrl: pq.StringArray{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
			"https://i.imgur.com/ExdKOOz.png",
			"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
	}
}
