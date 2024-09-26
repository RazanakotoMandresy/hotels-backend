package handler

import (
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
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
		uuids := uuid.New()
		res, err := s.services.Register(r.Context(), services.RegisterParams{
			UUID:     uuids,
			Name:     req.Name,
			Password: req.Passwords,
			Mail:     req.Mail,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " Services register error"}, http.StatusInternalServerError)
			return
		}
		tokenString, err := middleware.TokenManage(req.Mail, uuids.String())
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " Token's creation problem "}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseUsers{Users: *res, ResString: tokenString}, http.StatusOK)
	}
}
