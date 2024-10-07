package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s service) UploadImages() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
			files, handlers, err := r.FormFile("file")
			if err != nil {
				s.respond(w, errorResponse{Err: err.Error()}, http.StatusBadRequest)
				return
			}
			defer files.Close()
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{" valid uuid must provide in path"}, http.StatusBadRequest)
			return
		}

		res, err := s.services.UploadImages(r.Context(), uuid, files, handlers)
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " upload service"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseString{Res: res}, http.StatusOK)
	}
}
