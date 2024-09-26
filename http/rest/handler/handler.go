package handler

import (
	"github.com/RazanakotoMandresy/hotels-backend/internal/repository"
	services "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger *logrus.Logger
	services services.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger: lg,
		services: services.NewService(repository.NewRepository(db)),
	}
}
