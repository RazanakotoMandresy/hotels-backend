package handler

import (
	"net/http"

	// "github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func RegisterRoutes(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// adding logger middleware
	r.Use(handler.MiddlewareLogger(), middleware.AuthMiddleware)
	r.HandleFunc("/hotels", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/hotels", handler.GetAll()).Methods(http.MethodGet)
	r.HandleFunc("/hotels/{uuid}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/hotels/restore/{uuid}", handler.RestoreDeleted()).Methods(http.MethodPut)
	r.HandleFunc("/hotels/{uuid}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/hotels/{uuid}", handler.Delete()).Methods(http.MethodDelete)
	r.HandleFunc("/users/register", handler.Register()).Methods(http.MethodPost)
	r.HandleFunc("/users/login", handler.Login()).Methods(http.MethodPost)
}
