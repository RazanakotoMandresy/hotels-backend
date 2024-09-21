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
	Name        string       `valid:"required"`
	Description string       `valid:"required"`
	Services    pq.StringArray      
	Prix        uint         `valid:"required"`
	Status      model.Status `valid:"required"`
	Ouverture   string       `valid:"required"`
	// CreatedAt time.Time
}

func (s Service) Create(ctx context.Context, params CreateParams) (uuid.UUID, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return uuid.Nil, err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Hotels{
		UUID:        uuid.New(),
		Name:        params.Name,
		Description: params.Description,
		Services:    params.Services,
		Status:      params.Status,
		CreatedAt:   time.Now().UTC(),
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return uuid.Nil, err

	}

	err = tx.Commit()
	return entity.UUID, err
}
