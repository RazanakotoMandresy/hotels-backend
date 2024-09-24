package hotelsservices

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (s Service) GetAllHotels(ctx context.Context) (*[]model.Hotels, error) {
	hotels, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return &hotels, nil
}
