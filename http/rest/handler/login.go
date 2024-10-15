package handler

import (
	"net/http"

	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
)

func (s service) Login() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := new(userReq)
		if err := s.decode(r, req); err != nil {
			s.respond(w, errorResponse{err.Error() + " Decode's problems"}, http.StatusBadRequest)
			return
		}
		res, err := s.services.Login(r.Context(), services.LoginParams{
			Mail:     req.Mail,
			Password: req.Passwords,
		})
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " Services register error"}, http.StatusInternalServerError)
			return
		}
		tokenString, err := middleware.CreateToken(res.UUID.String(), res.Mail)
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " token's creation's error"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseUsers{Users: res, ResString: tokenString}, http.StatusOK)
	}
}
