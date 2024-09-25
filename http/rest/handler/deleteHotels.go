package handler

import (
	"fmt"
	"net/http"

)

func (s service) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := getUUUIDVars(r)
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, http.StatusBadRequest)
			return
		}
		// hotelsservices
		// hotelsservices
		if err := s.services.Delete(r.Context(),uuid); err != nil {
			if err.Error() == "sql: no rows in result set" {
				s.respond(w, errorResponse{fmt.Sprintf("the hotels with uuid %v has been deleted, you can restore it", uuid)}, http.StatusBadRequest)
				return
			}
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}

		s.respond(w, responseString{fmt.Sprintf("hotels with uuid : %v was successfuly deleted ", uuid)}, http.StatusOK)
	}
}
