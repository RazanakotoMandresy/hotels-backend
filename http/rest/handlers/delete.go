package handlers

import (
    "github.com/gorilla/mux"
    "net/http"
)

func (s service) Delete() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        uuid, err := vars["uuid"]
        if err  {
            s.respond(w, err, 0)
            return
        }

        errs := s.hotelsService.Delete(r.Context(), uuid)
        if errs != nil {
            s.respond(w, err, 0)
            return
        }
        s.respond(w, nil, http.StatusOK)
    }
}