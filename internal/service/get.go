package service

import (
	"context"
	"errors"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	// "github.com/RazanakotoMandresy/hotels-backend/pkg/db"
)

func (s Service) Get(ctx context.Context, uuid string) (model.Hotels, error) {
	hotels, err := s.repo.Find(ctx, uuid)
	switch {
	case err == nil:
	case errors.As(err, 404):
		return model.Hotels{}, errors.New("hotels object not found")
	default:
		return model.Hotels{}, err
	}
	return hotels, nil
}
