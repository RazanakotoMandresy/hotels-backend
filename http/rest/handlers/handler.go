package handlers

import (
	hotelsRepo "github.com/RazanakotoMandresy/deliveryapp-backend/internal/repository"
	hotelsService "github.com/RazanakotoMandresy/deliveryapp-backend/internal/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger        *logrus.Logger
	router        *mux.Router
	hotelsService hotelsService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:        lg,
		hotelsService: hotelsService.NewService(hotelsRepo.NewRepository(db)),
	}
}
