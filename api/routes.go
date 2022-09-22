package api

import "net/http"

func (s *Server) routes() {
	getR := s.router.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/posts", s.handleGreeter())
	postR := s.router.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/posts", s.handleGreeter())
}
