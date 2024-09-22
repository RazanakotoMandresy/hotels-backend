package handlers

import (
	"net/http"

	hotelsService "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/gorilla/mux"
)

func (s service) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {

			s.respond(w, errorResponse{"valid id must provide in path"}, http.StatusBadRequest)
			return
		}

		req := new(fullRequest)
		if err := s.decode(r, &req); err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}
		res, err := s.hotelsService.Update(r.Context(), hotelsService.UpdateParams{
			UUID:        uuid,
			Name:        &req.Name,
			Description: &req.Description,
			Prix:        &req.Prix,
			Status:      &req.Status,
			Ouverture:   &req.Ouverture,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}
		s.respond(w, modelResponse{res}, http.StatusOK)
	}
}
