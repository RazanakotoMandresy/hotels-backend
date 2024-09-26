package service

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (s Service) Get(ctx context.Context, uuid string) (*model.Hotels, error) {
	hotels, err := s.repo.Find(ctx, uuid)
	if err != nil {
		return nil, err
	}
	
	return hotels, nil
}
