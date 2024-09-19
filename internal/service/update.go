package service

import (
	"context"

	"github.com/RazanakotoMandresy/deliveryapp-backend/internal/model"
	"github.com/asaskevich/govalidator"
)

type UpdateParams struct {
	UUID        string `valid:"required"`
	Name        *string
	Description *string
	Status      *model.Status
}

func (s Service) Update(ctx context.Context, params UpdateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		// return erru.ErrArgument{Wrapped: err}
		return err
	}

	// find todo object
	todo, err := s.Get(ctx, params.UUID)
	if err != nil {
		return err
	}

	if params.Name != nil {
		todo.Name = *params.Name
	}
	if params.Description != nil {
		todo.Description = *params.Description
	}
	if params.Status != nil {
		if !params.Status.IsValid() {
			// return erru.ErrArgument{Wrapped: errors.New("given status not valid")}
			return err
		}
		todo.Status = *params.Status
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = s.repo.Update(ctx, todo)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
