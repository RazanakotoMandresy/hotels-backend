package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/pkg/db"
)

func (s Service) Get(ctx context.Context, uuid string) (model.Hotels, error) {
    fmt.Println("getServices")
	hotels, err := s.repo.Find(ctx, uuid)
	switch {
	case err == nil:
		fmt.Println("service1", err)
	case errors.As(err, &db.ErrObjectNotFound{}):
		fmt.Println("service2", err)
		return model.Hotels{}, errors.New("hotels object not found")
	default:
		fmt.Println("service3", err)
		return model.Hotels{}, err
	}
	fmt.Println("service4", err)
	return hotels, nil
}
