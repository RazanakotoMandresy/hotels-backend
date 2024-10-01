package service

import (
	"time"
	"fmt"
	"context"
	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
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


func (s Service) Create(ctx context.Context, params CreateParams) (*model.Hotels, error) {
	// just for oauth2 can be imported
	fmt.Println(oauth2.AccessTypeOffline)
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	userUUID := s.getUserUUIDInAuth(ctx)
	entity := model.Hotels{
		UUID:        uuid.New(),
		Name:        params.Name,
		Description: params.Description,
		Services:    params.Services,
		Status:      params.Status,
		Prix:        params.Prix,
		CreatedBy:  userUUID  ,
		CreatedAt:   time.Now().UTC(),
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return nil, err

	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
