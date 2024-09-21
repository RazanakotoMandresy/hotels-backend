package service

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/asaskevich/govalidator"
)

// no need services here
type UpdateParams struct {
	UUID        string `valid:"required"`
	Name        *string
	Description *string
	Prix        *uint
	Status      *model.Status
	Ouverture   *string
}

func (s Service) Update(ctx context.Context, params UpdateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		// return erru.ErrArgument{Wrapped: err}
		return err
	}
	// find hotels object
	hotels, err := s.Get(ctx, params.UUID)
	if err != nil {
		return err
	}

	if params.Name != nil {
		hotels.Name = *params.Name
	}
	if params.Description != nil {
		hotels.Description = *params.Description
	}
	if params.Status != nil {
		if !params.Status.IsValid() {
			return err
		}
		hotels.Status = *params.Status
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = s.repo.Update(ctx, hotels)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
