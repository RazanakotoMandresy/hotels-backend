package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s service) UploadImages() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{" valid uuid must provide in path"}, http.StatusBadRequest)
			return
		}
		// req := new(fileNameReq)
		// if err := s.decode(r, &req); err != nil {
		// 	s.respond(w, errorResponse{err.Error() + "decod's problemes"}, http.StatusInternalServerError)
		// 	return
		// }
		// w.Header().Set("Content-Type", "multipart/form-data")
		res, err := s.services.UploadImages(r.Context(), uuid, r)
		if err != nil {
			s.responseFormData(w, errorResponse{err.Error() + " upload service"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseString{Res: res}, http.StatusOK)
	}
}
