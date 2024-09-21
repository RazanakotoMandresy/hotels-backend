package handlers

import (
	"net/http"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	hotelsService "github.com/RazanakotoMandresy/hotels-backend/internal/service"

	"github.com/gorilla/mux"
)

func (s service) Update() http.HandlerFunc {
	type request struct {
		Name        *string       `json:"name"`
		Description *string       `json:"description"`
		Status      *model.Status `json:"status"`
	}

	type response struct {
		UUID string `json:"uuid"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {

			s.respond(w, errorResponse{"valid id must provide in path"}, 0)
			return
		}

		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		if err := s.decode(r, &req); err != nil {
			s.respond(w, errorResponse{err.Error()}, 0)
			return
		}
		err := s.hotelsService.Update(r.Context(), hotelsService.UpdateParams{
			UUID:        uuid,
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, 0)
			return
		}
		s.respond(w, response{UUID: uuid}, http.StatusOK)
	}
}
