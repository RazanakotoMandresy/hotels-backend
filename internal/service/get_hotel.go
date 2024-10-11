package service

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)


func (s Service) GetHotel(ctx context.Context, uuidOrName string) (*model.Hotels, error) {
	hotels, err := s.repo.Find(ctx, uuidOrName)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}
