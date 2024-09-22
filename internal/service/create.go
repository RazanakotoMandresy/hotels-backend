package service

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"time"
)

type CreateParams struct {
	Name        string `valid:"required"`
	Description string `valid:"required"`
	Services    pq.StringArray
	Prix        uint   `valid:"required"`
	Status      int    `valid:"required"`
	Ouverture   string `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (model.Hotels, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.Hotels{}, err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Hotels{}, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Hotels{
		UUID:        uuid.New(),
		Name:        params.Name,
		Description: params.Description,
		Services:    params.Services,
		Status:      params.Status,
		Prix:        params.Prix,
		CreatedAt:   time.Now().UTC(),
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return model.Hotels{}, err

	}

	err = tx.Commit()
	return entity, err
}
