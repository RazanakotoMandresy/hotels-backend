package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type responseString struct {
	res string `json:"res"`
}

func (s service) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{"veillez ajouter un uuid valide"}, 0)
			return
		}

		err := s.hotelsService.Delete(r.Context(), uuid)
		if err != nil {
			s.respond(w, errorResponse{err.Error()}, 0)
			return
		}
		s.respond(w, responseString{"effacer avec success"}, http.StatusOK)
	}
}
