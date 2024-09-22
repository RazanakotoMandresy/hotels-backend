package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s service) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]

		if !exist {
			s.respond(w, errorResponse{"no valid uuid in vars"}, http.StatusBadRequest)
			return
		}
		err := s.hotelsService.Delete(r.Context(), uuid)

		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				s.respond(w, errorResponse{fmt.Sprintf("the hotels with uuid %v has been deleted, you can restore it", uuid)}, http.StatusInternalServerError)
				return
			}
			s.respond(w, errorResponse{err.Error()}, http.StatusInternalServerError)
			return
		}

		s.respond(w, responseString{fmt.Sprintf("hotels with uuid : %v was successfuly deleted ", uuid)}, http.StatusOK)
	}
}
