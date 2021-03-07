package server

import (
	"Advertising/internal/app/api/utils"
	"Advertising/internal/app/model"
	"encoding/json"
	"net/http"
	"time"
)


func (s *server) handleGetAd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logRequest(r)
		ClintPage, err := utils.ParseUrlQueryID(r)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		result, err := s.store.Advertisement().GetOne(ClintPage)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, result)

	}
}

func (s *server) handleGetAds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logRequest(r)
		ClintPage, err := utils.ParseUrlQueryPage(r)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		result, err := s.store.Advertisement().GetAllMap(ClintPage)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, result)

	}
}

func (s *server) handleCreateAd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logRequest(r)

		a := &model.Advertisement{}
		if err := json.NewDecoder(r.Body).Decode(a); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		a.Date = time.Now()
		id, err := s.store.Advertisement().Create(a)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, map[string]int{"id": id})
	}
}
