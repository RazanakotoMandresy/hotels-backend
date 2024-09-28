package handler

import (
	"net/http"

	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

//	func authentifiedRoutes(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
//		handler := newHandler(lg, db)
//		// adding logger middleware
//		r.Use(middleware.AuthMiddleware, handler.MiddlewareLogger())
//	}
func routes(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// no need to be authentified
	r.Use(handler.MiddlewareLogger())
	r.HandleFunc("/hotels", handler.GetAll()).Methods(http.MethodGet)
	r.HandleFunc("/hotels/{uuid}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/users/register", handler.Register()).Methods(http.MethodPost)
	r.HandleFunc("/users/login", handler.Login()).Methods(http.MethodPost)
	private := r.PathPrefix("/").Subrouter()
	private.Use(middleware.AuthMiddleware, handler.MiddlewareLogger())
	private.HandleFunc("/hotels", handler.Create()).Methods(http.MethodPost)
	private.HandleFunc("/hotels/{uuid}", handler.Update()).Methods(http.MethodPut)
	private.HandleFunc("/hotels/restore/{uuid}", handler.RestoreDeleted()).Methods(http.MethodPut)
	private.HandleFunc("/hotels/{uuid}", handler.Delete()).Methods(http.MethodDelete)

}
