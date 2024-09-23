package handlers

import (
	"net/http"

	hotelsService "github.com/RazanakotoMandresy/hotels-backend/internal/service"
)

func (s service) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(fullRequest)
		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " decode's problems"}, http.StatusInternalServerError)
			return
		}

		res, err := s.hotelsService.Create(r.Context(), hotelsService.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
			Ouverture:   req.Ouverture,
			Prix:        req.Prix,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " error on happen on the create handler from services"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, modelResponse{res}, http.StatusCreated)
	}
}
