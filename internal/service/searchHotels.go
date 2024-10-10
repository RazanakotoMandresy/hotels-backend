package service

import (
	"github.com/RazanakotoMandresy/hotels-backend/internal/model"

	"context"
)

func (s Service) SearchHotels(ctx context.Context, searched string) ([]model.Hotels, error) {
	usersFound, err := s.repo.SearchQuery(ctx, searched)
	if err != nil {
		return nil, err
	}

	return usersFound, nil
}
