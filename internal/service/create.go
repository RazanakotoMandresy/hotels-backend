package service

import (
	"context"
	"fmt"

	"github.com/RazanakotoMandresy/deliveryapp-backend/internal/model"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"

	"time"
)

type CreateParams struct {
	Name        string       `valid:"required"`
	Description string       `valid:"required"`
	Status      model.Status `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (string, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return "", err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return "", err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Hotels{
		UUID:        uuid.NewString(),
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
		CreatedOn:   time.Now().UTC(),
	}
	fmt.Println(entity.UUID)
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		fmt.Println("err services ", err)
		return "", err
	}

	err = tx.Commit()
	return entity.UUID, err
}
