package service

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
)

// TODO implements the payments
func (s Service) ReserveHotel(ctx context.Context, uuidHotels string, params ReserveParams) (*model.Hotels, error) {
	hotels, err := s.GetHotel(ctx, uuidHotels)
	Date := params.Starting_date + "->" + params.Ending_date
	if err != nil {
		return nil, err
	}
	uuidUsr := middleware.GetUserUUIDInAuth(ctx)
	user, err := s.repo.FindUserByUUID(ctx, uuidUsr)
	if err != nil {
		return nil, err
	}
	decriptedPassword, err := middleware.Decrypt(user.Passwords)
	if err != nil {
		return nil, err
	}
	if decriptedPassword != params.Password {
		return nil, errors.New(" wrong passwords ")
	}
	if !hotels.Status {
		return nil, fmt.Errorf("hotels %v is not available", hotels.Name)
	}
	if slices.Contains(hotels.ReservationLists, Date) {
		return nil, fmt.Errorf("hotels %v already reserved on %v", hotels.Name, Date)
	}
	hotels.ReservationLists = append(hotels.ReservationLists, Date)
	if err := s.repo.UpdateHotel(ctx, *hotels); err != nil {
		return nil, err
	}
	
	return hotels, nil
}
