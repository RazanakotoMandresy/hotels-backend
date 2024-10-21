package service

import (
	"context"
	"errors"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/asaskevich/govalidator"
)

// type resFilter struct {
// 	Name      string
// 	Ouverture string
// 	Place     string
// 	Services  string
// 	Prix      uint
// }

func (s Service) FilterHotels(ctx context.Context, params FilterParams) ([][]model.Hotels, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}
	userUUID := middleware.GetUserUUIDInAuth(ctx)
	if userUUID == "" {
		return nil, errors.New("no uuid in bearer auth")
	}
	res, err := params.checkParams(ctx, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (p FilterParams) checkParams(ctx context.Context, s Service) ([][]model.Hotels, error) {
	var arrOfarrHotels = [][]model.Hotels{}
	if p.Name != "" {
		hotels, err := s.repo.Filter(ctx, p.Name, "name")
		if err != nil {
			return nil, err
		}
		arrOfarrHotels = append(arrOfarrHotels, hotels)
	}
	if p.Place != "" {
		hotels, err := s.repo.Filter(ctx, p.Place, "place")
		if err != nil {
			return nil, err
		}
		arrOfarrHotels = append(arrOfarrHotels, hotels)
	}
	if p.Ouverture != "" {
		hotels, err := s.repo.Filter(ctx, p.Ouverture, "ouverture")
		if err != nil {
			return nil, err
		}
		arrOfarrHotels = append(arrOfarrHotels, hotels)
	}
	return arrOfarrHotels, nil
}
