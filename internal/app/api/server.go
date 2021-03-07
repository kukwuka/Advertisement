package server

import (
	"Advertising/internal/app/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	store  store.Store
}

//New server
func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s

}

//Server
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

//error
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	logrus.Println(err)
	s.respond(w, r, code, map[string]string{"error": err.Error()})

}

//respond
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

//logRequest
func (s *server) logRequest(r *http.Request) {
	logrus.Printf("request to:%s , from: %s ", r.URL, r.Host)
}
