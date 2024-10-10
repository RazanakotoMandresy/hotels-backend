package handler

import "net/http"

func (s service) SearchHotels() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("hotels")
		if search == "" {
			s.respond(w, responseString{"please enter hotel's name to search"}, http.StatusBadRequest)
			return
		}
		
	}
}
