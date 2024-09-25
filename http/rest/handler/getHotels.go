package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s service) Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		numUUID, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{Err: "valid uuid must provide in path"}, http.StatusBadRequest)
			return
		}
		getResponse, err := s.services.Get(r.Context(), numUUID)
		if err != nil {
			s.respond(w, errorResponse{Err: err.Error() + " error on get services"}, http.StatusNotFound)
			return
		}
		s.respond(w, modelResponse{Hotels: *getResponse}, http.StatusOK)
	}
}
