package hotelshandlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	// adding logger middleware
	r.Use(handler.MiddlewareLogger())
	r.HandleFunc("/hotels", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/hotels/{uuid}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/hotels/restore/{uuid}", handler.RestoreDeleted()).Methods(http.MethodPut)
	r.HandleFunc("/hotels/{uuid}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/hotels/{uuid}", handler.Delete()).Methods(http.MethodDelete)
}
