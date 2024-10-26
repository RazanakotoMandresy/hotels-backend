package service

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/google/uuid"
)

// TODO implements the payments
func (s Service) ReserveHotel(ctx context.Context, uuidHotels string, params ReserveParams) (*model.Hotels, error) {
	uuidUsr := middleware.GetUserUUIDInAuth(ctx)
	Date := params.Starting_date + "->" + params.Ending_date
	hotels, err := s.repo.FindHotel(ctx, uuidHotels)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindUserByUUID(ctx, uuidUsr)
	if err != nil {
		return nil, err
	}
	decriptedPassword, err := middleware.Decrypt(user.Passwords)
	if err != nil {
		return nil, err
	}
	if err := validDate(params); err != nil {
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
	if err := s.repo.CreateReservation(ctx, &model.Reservation{
		UUID:              uuid.New(),
		ReservedBy:        uuidUsr,
		ReservationStart:  params.Starting_date,
		ReservationEnding: params.Ending_date,
	}); err != nil {
		return nil, err
	}
	return hotels, nil
}

func validDate(r ReserveParams) error {
	splitedStart := strings.Split(r.Starting_date, "-")
	splitedEnding := strings.Split(r.Ending_date, "-")
	if len(splitedEnding) != 3 || len(splitedStart) != 3 {
		return fmt.Errorf("invalid time format should be like 2025-01-02 yyyy-mm-dd yours : starting : %v,ending %v", splitedStart, splitedEnding)
	}
	dateStart, err := convInt(splitedStart)
	if err != nil {
		return err
	}
	dateEnd, err := convInt(splitedEnding)
	if err != nil {
		return err
	}

	start := time.Date(dateStart[0], time.Month(dateStart[1]), dateStart[2], 0, 0, 0, 0, time.Local)
	end := time.Date(dateEnd[0], time.Month(dateEnd[1]), dateStart[2], 0, 0, 0, 0, time.Local)
	// year date[0] month date[1] day[2]
	if time.Now().Compare(start) == +1 {
		return fmt.Errorf(" cannot reserve in a past date ")
	}
	if time.Time.Compare(start, end) == +1 {
		return fmt.Errorf(" end date is before start : %v and end %v", start, end)
	}
	return nil
}
func convInt(dateArray []string) (map[int]int, error) {
	date := make(map[int]int)
	for n, value := range dateArray {
		int, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("all date should be an number :%v ", err)
		}
		date[n] = int
	}
	return date, nil
}
