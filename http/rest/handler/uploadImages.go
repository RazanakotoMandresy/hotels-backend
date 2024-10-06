package handler

import (
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
)

func (s service) UploadImages() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		file := make(chan multipart.File)
		handler := make(chan *multipart.FileHeader)
		go func() {
			files, handlers, err := r.FormFile("file")
			if err != nil {
				s.respond(w, errorResponse{Err: err.Error()}, http.StatusBadRequest)
				return
			}
			file <- files
			handler <- handlers
			defer files.Close()
		}()
		vars := mux.Vars(r)
		uuid, exist := vars["uuid"]
		if !exist {
			s.respond(w, errorResponse{" valid uuid must provide in path"}, http.StatusBadRequest)
			return
		}
		res, err := s.services.UploadImages(r.Context(), uuid, file, handler)
		if err != nil {
			s.respond(w, errorResponse{err.Error() + " upload service"}, http.StatusInternalServerError)
			return
		}
		s.respond(w, responseString{Res: res}, http.StatusOK)
	}
}
