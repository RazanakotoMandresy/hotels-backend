package handler

import (
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	// "github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/google/uuid"

	"net/http"
)

func (s service) Register() http.HandlerFunc {
	type registerReq struct {
		Name      string `json:"name"`
		Passwords string `json:"passwords"`
		Mail      string `json:"mail"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(registerReq)
		if err := s.decode(r, req); err != nil {
			s.respond(w, errorResponse{err.Error() + " Decode's problems"}, http.StatusBadRequest)
			return
		}
		if len(req.Passwords) < 8 {
			s.respond(w, errorResponse{"password too short "}, http.StatusBadRequest)
			return
		}
		userUUUID := uuid.New()
		res, err := s.services.Register(r.Context(), services.RegisterParams{
			UUID:     userUUUID,
			Name:     req.Name,
			Password: req.Passwords,
			Mail:     req.Mail,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " Services register error"}, http.StatusInternalServerError)
			return
		}
		tokenString, err := middleware.CreateToken(userUUUID.String(), req.Mail)
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " token's creation"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseUsers{Users: res, ResString: tokenString}, http.StatusOK)
	}
}
