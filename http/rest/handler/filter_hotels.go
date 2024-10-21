package handler

import (
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"net/http"
)

func (s service) filterHotels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(filterRequest)
		if err := s.decode(r, &req); err != nil {
			s.respond(w, errorResponse{err.Error() + " decode's error"}, http.StatusBadRequest)
			return
		}
		hotels, err := s.services.FilterHotels(r.Context(), services.FilterParams{
			Name:      req.Name,
			Prix:      req.Prix,
			Place:     req.Place,
			Service:   req.Service,
			Ouverture: req.Ouverture,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " services error "}, http.StatusInternalServerError)
			return
		}
		s.respond(w, hotels, http.StatusOK)
	}
}
