package handler

import (
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"net/http"
)

func (s service) Register() http.HandlerFunc {
	type registerReq struct {
		name      string `json:"name"`
		passwords string `json:"passwords"`
		mail      string `json:"mail"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(registerReq)
		if err := s.decode(r, req); err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusBadRequest)
			return
		}
		res, err := s.services.Register(r.Context(), services.RegisterParams{
			Name:     req.name,
			Password: req.passwords,
			Mail:     req.mail,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}
		s.respond(w, modelResponse{Users: *res}, http.StatusOK)
	}
}
