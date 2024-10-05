package handler

import (
	"net/http"

	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/gorilla/mux"
)

func (s service) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{" valid uuid must provide in path"}, http.StatusBadRequest)
			return
		}
		req := new(fullRequest)
		if err := s.decode(r, &req); err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}

		res, err := s.services.UpdateHotels(r.Context(), services.UpdateParams{
			UUID:        uuid,
			Name:        &req.Name,
			Description: &req.Description,
			Prix:        &req.Prix,
			Status:      &req.Status,
			Ouverture:   &req.Ouverture,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " update service"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responsesHotels{
			UUID:        res.UUID,
			Name:        res.Name,
			Description: res.Description,
			Status:      res.Status,
			Created_at:  res.CreatedAt,
			Prix:        res.Prix,
			Updated_at:  res.UpdatedAt,
		}, http.StatusOK)
	}
}
