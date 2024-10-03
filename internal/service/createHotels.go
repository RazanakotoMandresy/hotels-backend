package service

import (
	"context"
	"fmt"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/oauth2"
)

type CreateParams struct {
	Name        string `valid:"required"`
	Description string `valid:"required"`
	Services    pq.StringArray
	Prix        uint   `valid:"required"`
	Status      int    `valid:"required"`
	Ouverture   string `valid:"required"`
}

func (s Service) CreateHotel(ctx context.Context, params CreateParams) (*model.Hotels, error) {
	userUUID := middleware.GetUserUUIDInAuth(ctx)
	hotelsUUID := uuid.New()
	// just for oauth2 can be imported
	fmt.Println(oauth2.AccessTypeOffline)
	// find the user to be updated in his hotels list
	users, err := s.repo.FindUserByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	entity := model.Hotels{
		UUID:        hotelsUUID,
		Name:        params.Name,
		Description: params.Description,
		Services:    params.Services,
		Status:      params.Status,
		Prix:        params.Prix,
		CreatedBy:   userUUID,
		CreatedAt:   time.Now().UTC(),
	}
	if err := s.repo.Create(ctx, &entity); err != nil {
		return nil, err
	}
	users.ListHotels = append(users.ListHotels, hotelsUUID.String())
	if err := s.repo.UpdateUser(ctx, *users); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &entity, nil
}
