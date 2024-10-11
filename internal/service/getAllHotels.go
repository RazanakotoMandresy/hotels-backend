package service

import (
	"context"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (s Service) GetAllHotels(ctx context.Context) (*[]model.Hotels, error) {
	hotels, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second * 10)
	return &hotels, nil
}
