package handlers

import (
	"errors"

	"net/http"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/gorilla/mux"
)

func (s service) Get() http.HandlerFunc {
	type response struct {
		UUID        string       `json:"uuid"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		numUUID, exist := vars["uuid"]
		if !exist {
			s.respond(w, errors.New("valid uuid must provide in path"), 0)
			return
		}
		getResponse, err := s.hotelsService.Get(r.Context(), numUUID)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			UUID:        getResponse.UUID,
			Name:        getResponse.Name,
			Description: getResponse.Description,
			Status:      getResponse.Status,
			CreatedOn:   getResponse.CreatedOn,
			UpdatedOn:   getResponse.UpdatedOn,
		}, http.StatusOK)
	}
}
