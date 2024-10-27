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
	if decriptedPassword != params.Password {
		return nil, errors.New(" wrong passwords ")
	}
	if !hotels.Status {
		return nil, fmt.Errorf("hotels %v is not available", hotels.Name)
	}
	if err := validDate(params, hotels.ReservationLists); err != nil {
		return nil, err
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
func parseReservation(prevReservations []string, newStartDate, newEndDate *time.Time) error {
	for _, prevReservation := range prevReservations {
		dates := strings.Split(prevReservation, "->")
		if len(dates) != 2 {
			return errors.New("invalid date format")
		}
		startDate, err := time.Parse("2006-01-02", dates[0])
		if err != nil {
			return err
		}
		endDate, err := time.Parse("2006-01-02", dates[1])
		if err != nil {
			return err
		}
		if newStartDate.Before(startDate) && newEndDate.After(endDate) {
			return fmt.Errorf("cannot use date between %v and %v already taken until %v -> %v ", newStartDate, newEndDate, startDate, endDate)
		}
		valbool := newStartDate.Before(startDate) && newEndDate.After(endDate)
		fmt.Printf("prevdateStart %v , prevdateEnd %v , newstart %v , newend %v \n ", startDate, endDate, newStartDate, newEndDate)
		fmt.Printf("value bool %v ", valbool)
	}
	return nil
}
func strToTime(dateArray []string) (*time.Time, error) {
	date := make(map[int]int)
	for n, value := range dateArray {
		int, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("all date should be an number :%v ", err)
		}
		date[n] = int
	}
	datetTime := time.Date(date[0], time.Month(date[1]), date[2], 0, 0, 0, 0, time.UTC)
	return &datetTime, nil
}

// func parseReservation(date string) (err , )
func validDate(r ReserveParams, prevReserv []string) error {
	splitedStart := strings.Split(r.Starting_date, "-")
	splitedEnding := strings.Split(r.Ending_date, "-")
	if len(splitedEnding) != 3 || len(splitedStart) != 3 {
		return fmt.Errorf("invalid time format should be like 2025-01-02 yyyy-mm-dd yours : starting : %v,ending %v", splitedStart, splitedEnding)
	}
	dateStart, err := strToTime(splitedStart)
	if err != nil {
		return err
	}
	dateEnd, err := strToTime(splitedEnding)
	if err != nil {
		return err
	}
	// year date[0] month date[1] day[2]
	if time.Now().Compare(*dateStart) == +1 {
		return fmt.Errorf(" cannot reserve in a past date ")
	}
	if time.Time.Compare(*dateStart, *dateEnd) == +1 {
		return fmt.Errorf(" end date is before start : %v and end %v", dateStart, dateEnd)
	}
	if err := parseReservation(prevReserv, dateStart, dateEnd); err != nil {
		return err
	}
	return nil
}
