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

// type NilHotel model.Hotels

func (s Service) Update(ctx context.Context, params UpdateParams) (model.Hotels, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		// return erru.ErrArgument{Wrapped: err}
		return model.Hotels{}, err
	}
	// find hotels object
	hotels, err := s.Get(ctx, params.UUID)
	if err != nil {
		return model.Hotels{}, err
	}

	if params.Name != nil {
		hotels.Name = *params.Name
	}
	if params.Description != nil {
		hotels.Description = *params.Description
	}
	if params.Prix != nil {
		hotels.Prix = *params.Prix
	}
	if params.Status != nil {
		if !params.Status.IsValid() {
			return model.Hotels{}, err
		}
		hotels.Status = *params.Status
	}
	if params.Ouverture != nil {
		hotels.Ouverture = *params.Ouverture
	}
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Hotels{}, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = s.repo.Update(ctx, hotels)
	if err != nil {
		return model.Hotels{}, err
	}

	err = tx.Commit()
	return model.Hotels{}, err
}
