package service

import (
	"context"
	"errors"

	"github.com/RazanakotoMandresy/deliveryapp-backend/internal/model"
	"github.com/RazanakotoMandresy/deliveryapp-backend/pkg/db"
)
func (s Service) Get(ctx context.Context, uuid string) (model.Hotels, error) {
    todo, err := s.repo.Find(ctx, uuid)
    switch {
    case err == nil:
    case errors.As(err, &db.ErrObjectNotFound{}):
        return model.Hotels{}, errors.New("todo object not found")
        // return model.ToDo{}, erru.ErrArgument{errors.New("todo object not found")}
	default:
        return model.Hotels{}, err
    }
    return todo, nil
}