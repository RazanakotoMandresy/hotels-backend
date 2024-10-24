package handler

import (
	"net/http"
)

func (s service) SearchHotels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("hotels")
		if search == "" {
			s.respond(w, errorResponse{"please enter hotel's name to search"}, http.StatusBadRequest)
			return
		}
		hotels, err := s.services.SearchHotels(r.Context(), search)
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusNotFound)
			return
		}
		s.respond(w, arrayHotelsResponse{hotels}, http.StatusOK)
	}
}
