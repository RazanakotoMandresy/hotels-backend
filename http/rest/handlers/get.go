package handlers

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
		getResponse, err := s.hotelsService.Get(r.Context(), numUUID)
		if err != nil {
			s.respond(w, errorResponse{Err: err.Error() + "get services"}, http.StatusNotFound)
			return
		}
		s.respond(w, response{
			UUID:        getResponse.UUID,
			Name:        getResponse.Name,
			Description: getResponse.Description,
			Status:      getResponse.Status,
			Created_at:  getResponse.CreatedAt,
			Updated_at:  getResponse.UpdatedAt,
		}, http.StatusOK)
	}
}
