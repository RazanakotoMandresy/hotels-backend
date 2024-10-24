package handler

import (
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func (s service) UserReservation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{" valid uuid must be provide in path "}, http.StatusBadRequest)
			return
		}
		req := new(reserveRequests)
		hotels, err := s.services.ReserveHotel(r.Context(), uuid, services.ReserveParams{
			Date:     req.Date,
			Password: req.Password,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusBadRequest)
			return
		}
		s.respond(w, responseHotel{*hotels}, http.StatusOK)
	}
}
