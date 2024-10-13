package handler

import (
	"net/http"

	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
)

func (s service) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(fullRequest)
		if err := s.decode(r, &req); err != nil {
			s.respond(w, errorResponse{err.Error() + " decode's problems"}, http.StatusInternalServerError)
			return
		}
		res, err := s.services.CreateHotel(r.Context(), services.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
			Ouverture:   req.Ouverture,
			Prix:        req.Prix,
			Place:       req.Place,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " error on happen on the create handler from services"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, res, http.StatusCreated)
	}
}
