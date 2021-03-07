package server

import (
	"Advertising/internal/app/store/teststore"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handleCreateAd(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]interface{}{
				"name":"первое",
				"description": "описание",
				"cost":40,
				"img_url":[]string{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
				"https://i.imgur.com/ExdKOOz.png",
				"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "without name",
			payload: map[string]interface{}{
				"description": "описание",
				"cost":40,
				"img_url":[]string{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
					"https://i.imgur.com/ExdKOOz.png",
					"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "without cost",
			payload: map[string]interface{}{
				"name":"первое",
				"description": "описание",
				"img_url":[]string{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
					"https://i.imgur.com/ExdKOOz.png",
					"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid url of image",
			payload: map[string]interface{}{
				"name":"первое",
				"cost":40,
				"description": "описание",
				"img_url":[]string{"http://www.dsaolangprograms.com/skin/frontendsad/base/default/logo.png",
					"https://i.imdsagur.com/ExdKOOz.png",
					"http://site.medaishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_,code , err:= createAdForTest(tc.payload,s)
			if err!=nil{
				assert.Error(t, err)
			}
			assert.Equal(t, tc.expectedCode, code)
		})
	}
}

func TestServer_handleGetAds(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      map[string]string
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"page":"1",
			},
			expectedCode: http.StatusOK,
		},
		//{
		//	name:         "negative",
		//	payload: map[string]string{
		//		"page":"-1",
		//	},
		//	expectedCode: http.StatusUnprocessableEntity,
		//},
		//{
		//	name:         "out of range",
		//	payload: map[string]string{
		//		"page":"1000000",
		//	},
		//	expectedCode: http.StatusUnprocessableEntity,
		//},
	}
	//Create 1 for testing
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_,code , err:= createAdForTest(map[string]interface{}{
				"name":"первое",
				"description": "описание",
				"cost":40,
				"img_url":[]string{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
					"https://i.imgur.com/ExdKOOz.png",
					"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
			},s)
			if err != nil {
				assert.Error(t, err)
			}
			if code != http.StatusCreated {
				assert.Error(t, errors.New("not 201"))

			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/ads/", nil)
			q := req.URL.Query()
			q.Add("page",tc.payload["page"])
			s.ServeHTTP(rec, req)
			logrus.Print(rec.Body)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}
}

func TestServer_handleGetAd(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      map[string]string
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"page":"1",
			},
			expectedCode: http.StatusOK,
		},
	}
	//Create 1 for testing
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			id,code , err:= createAdForTest(map[string]interface{}{
				"name":"первое",
				"description": "описание",
				"cost":40,
				"img_url":[]string{"http://www.golangprograms.com/skin/frontend/base/default/logo.png",
					"https://i.imgur.com/ExdKOOz.png",
					"http://site.meishij.net/r/58/25/3568808/a3568808_142682562777944.jpg"},
			},s)
			if err != nil {
				assert.Error(t, err)
			}
			if code != http.StatusCreated {
				assert.Error(t, errors.New("not 201"))

			}

			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/ad/", nil)
			q := req.URL.Query()
			q.Add("id",fmt.Sprint(id))
			s.ServeHTTP(rec, req)
			log.Print(rec.Body)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}
}



func createAdForTest(payload interface{} , s *server) (Id int ,StatusCode int,err error){
	b := &bytes.Buffer{}
	err =json.NewEncoder(b).Encode(payload)
	if err != nil {
		return 0,0 ,err
	}
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/ad/", b)
	s.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated{
		return 0,rec.Code, errors.New("didn't created")
	}
	var m map[string]int
	err = json.Unmarshal(rec.Body.Bytes() ,&m)
	if err != nil {
		return 0,0,err
	}
	return 0,rec.Code,nil
}