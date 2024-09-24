package hotelshandlers

import (
	"github.com/RazanakotoMandresy/hotels-backend/internal/repository"
	hotelsservices "github.com/RazanakotoMandresy/hotels-backend/internal/service/hotelsServices"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger        *logrus.Logger
	hotelsService hotelsservices.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:        lg,
		hotelsService: hotelsservices.NewService(repository.NewRepository(db)),
	}
}
