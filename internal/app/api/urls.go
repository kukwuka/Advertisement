package server

func (s *server) configureRouter() {
	s.router.HandleFunc("/ad/", s.handleGetAd()).Methods("GET")
	s.router.HandleFunc("/ads/", s.handleGetAds()).Methods("GET")
	s.router.HandleFunc("/ad/", s.handleCreateAd()).Methods("POST")
}
