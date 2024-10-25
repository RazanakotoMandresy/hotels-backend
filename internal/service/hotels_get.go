package service

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (s Service) GetHotel(ctx context.Context, uuid string) (*model.Hotels, error) {
	hotels, err := s.repo.FindHotel(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
