package handlers

import (
	"errors"

	"net/http"
	"strconv"
	"time"

	"github.com/RazanakotoMandresy/deliveryapp-backend/internal/model"
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

		numUUID, err := strconv.Atoi(vars["uuid"])
		if err != nil {
			// s.respond(w, erru.ErrArgument{
			//     Wrapped: errors.New("valid id must provide in path"),
			// }, 0)
			s.respond(w, errors.New("valid id must provide in path"), 0)
			return
		}

		getResponse, err := s.hotelsService.Get(r.Context(), string(numUUID))
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
