package model_test

import (
	"Advertising/internal/app/model"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testsPositiveCases = []struct {
	name    string
	value   float64
	isValid bool
}{
	{"Negative",
		-3.2,
		false,
	},
	{"Zero",
		0,
		true,
	},
	{"Positive",
		23.31,
		true,
	},
}

var testValidateImg = []struct {
	name    string
	value   interface{}
	isValid bool
}{
	{"3 Url",
		pq.StringArray{
			"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
			"https://i.imgur.com/ExdKOOz.png",
			"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg",
		},
		true,
	},
	{"1 Url",
		pq.StringArray{
			"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
		},
		true,
	},
	{"3 Incorrect Url",
		pq.StringArray{
			"http://www.golangpdasdadrograms.com/skin/frontend/base/default/logo.png",
			"https://i.imdsadasgfdsfdsfdsfdsur.com/ExdKOOz.pnfdsg",
			"http://site.medasishij.net/r/58/25/3568808/a3568808_142682562777944.jpg",
		},
		false,
	},
	{"Smt nill",
		pq.StringArray{
			"adsads",
			"",
			"",
		},
		false,
	},
}

func TestIsPositive(t *testing.T) {

	for _, tc := range testsPositiveCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, model.IsPositive(tc.value))
			} else {
				assert.Error(t, model.IsPositive(tc.value))
			}
		})
	}
}

func TestValidateImg(t *testing.T) {
	for _, tc := range testValidateImg {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, model.ValidateImg(tc.value))
			} else {
				assert.Error(t, model.ValidateImg(tc.value))
			}
		})
	}
}

func TestAllValidationInternal(t *testing.T) {
	//assert.NoError(t, nil)
	TestIsPositive(t)
	TestValidateImg(t)

}
