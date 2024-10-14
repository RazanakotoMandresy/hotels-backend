package handler

import (
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"net/http"
)

// filtre les hotels par primo , lieux , prix , de mivoka array amzay le izy apres
// TODO search an json array of all of the most visited and toristic place in the word
func (s service) filterHotels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(filterRequest)
		if err := s.decode(r, &req); err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusBadRequest)
			return
		}
		hotels, err := s.services.FilterHotels(r.Context(), services.FilterParams{})
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}
		s.respond(w, arrayHotelsResponse{hotels}, http.StatusOK)
	}
}
