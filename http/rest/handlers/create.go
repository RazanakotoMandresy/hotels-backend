package handlers

import (
	"net/http"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	hotelsService "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/google/uuid"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
	}

	type response struct {
		UUID uuid.UUID `json:"uuid"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		
		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, errorResponse{err.Error() + "decode's problems"}, 500)
			return
		}

		uuid, err := s.hotelsService.Create(r.Context(), hotelsService.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}
		s.respond(w, response{UUID: uuid}, http.StatusOK)
	}
}
