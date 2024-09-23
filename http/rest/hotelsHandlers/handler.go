package hotelshandlers

import (
	hotelsRepo "github.com/RazanakotoMandresy/hotels-backend/internal/repository"
	hotelsService "github.com/RazanakotoMandresy/hotels-backend/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger        *logrus.Logger
	hotelsService hotelsService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:        lg,
		hotelsService: hotelsService.NewService(hotelsRepo.NewRepository(db)),
	}
}
