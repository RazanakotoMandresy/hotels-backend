package service

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/asaskevich/govalidator"
)

type filterResHotels struct {
	MatchedName      []model.Hotels `json:"matched_name"`
	MatchedPlace     []model.Hotels `json:"matched_place"`
	MatchedOuverture []model.Hotels `json:"matched_ouverture"`
	MatchedServie    []model.Hotels `json:"matched_service"`
	MatchedPrice     []model.Hotels `json:"matched_price"`
}

func (s Service) FilterHotels(ctx context.Context, params FilterParams) ([]filterResHotels, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}
	res, err := params.checkParams(ctx, s)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (f FilterParams) checkParams(ctx context.Context, s Service) ([]filterResHotels, error) {
	var arrFiltedHotel []filterResHotels
	if f.Name != "" {
		hotels, err := s.repo.Filter(ctx, f.Name, "name")
		if err != nil {
			return nil, err
		}
		arrFiltedHotel = append(arrFiltedHotel, filterResHotels{MatchedName: hotels})
	}
	if f.Place != "" {
		hotels, err := s.repo.Filter(ctx, f.Place, "place")
		if err != nil {
			return nil, err
		}
		arrFiltedHotel = append(arrFiltedHotel, filterResHotels{MatchedPlace: hotels})
	}
	if f.Ouverture != "" {
		hotels, err := s.repo.Filter(ctx, f.Ouverture, "ouverture")
		if err != nil {
			return nil, err
		}
		arrFiltedHotel = append(arrFiltedHotel, filterResHotels{MatchedOuverture: hotels})
	}
	// check if min and max are both not equal to zero
	if (f.MinBudget * f.MaxBudget) != 0 {
		// handle the price matching
		hotels, err := s.repo.FilterPrice(ctx, f.MinBudget, f.MaxBudget)
		if err != nil {
			return nil, err
		}
		arrFiltedHotel = append(arrFiltedHotel, filterResHotels{MatchedPrice: hotels})
	}
	// if f.Service != "" {
	// }
	return arrFiltedHotel, nil
}
