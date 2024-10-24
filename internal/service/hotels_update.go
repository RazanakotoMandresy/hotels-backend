package service

import (
	"context"
	"errors"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/asaskevich/govalidator"
)

// no need services here

func (s Service) UpdateHotels(ctx context.Context, params UpdateParams) (*model.Hotels, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}
	// find hotels object
	hotels, err := s.GetHotel(ctx, params.UUID)
	if err != nil {
		return nil, err
	}
	userUUID := middleware.GetUserUUIDInAuth(ctx)
	if userUUID == "" {
		return nil, errors.New("no uuid in bearer auth")
	}
	if userUUID != hotels.CreatedBy {
		return nil, errors.New("you are not the creator of this hotels")
	}
	if params.Name != nil {
		hotels.Name = *params.Name
	}
	if params.Description != nil {
		hotels.Description = *params.Description
	}
	if params.Prix != nil {
		hotels.Prix = *params.Prix
	}
	if params.Status != nil {
		hotels.Status = *params.Status
	}
	if params.Ouverture != nil {
		hotels.Ouverture = *params.Ouverture
	}
	now := time.Now().UTC()
	hotels.UpdatedAt = &now
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	if err := s.repo.UpdateHotel(ctx, *hotels); err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return hotels, err
	}
	return hotels, nil
}
