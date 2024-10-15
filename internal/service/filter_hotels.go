package service

import (
	"context"
	"errors"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/asaskevich/govalidator"
)

func (s Service) FilterHotels(ctx context.Context, params FilterParams) ([]model.Hotels, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}
	userUUID := middleware.GetUserUUIDInAuth(ctx)
	if userUUID == "" {
		return nil, errors.New("no uuid in bearer auth")
	}
	hotelsFilterd, err := s.repo.FilterHotels(ctx, params.Name, params.Ouverture, params.Place, params.Service, params.Prix)
	if err != nil {
		return nil, err
	}
	return hotelsFilterd, nil
}
