package hotelshandlers

import "net/http"

func (s service) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hotels, err := s.hotelsService.GetAllHotels(r.Context())
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}
		s.respond(w, arrayHotels{*hotels}, http.StatusOK)
	}
}
