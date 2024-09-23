package handlers

import (
	"fmt"
	"net/http"
)

func (s service) RestoreDeleted() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := getUUUIDVars(r)
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusBadRequest)
			return
		}
		if err := s.hotelsService.RestoreDeleted(r.Context(), uuid); err != nil {
			s.respond(w, errorResponse{err.Error() + " Error services Retored on handler"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseString{fmt.Sprintf("hotels with uuid %v restored perfectly", uuid)}, http.StatusOK)
	}
}
