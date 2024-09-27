package service

import (
	"context"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/google/uuid"
)

type RegisterParams struct {
	Name     string `valid:"required"`
	Mail     string `valid:"required"`
	Password string `valid:"required"`
	UUID     uuid.UUID
}

func (s Service) Register(ctx context.Context, params RegisterParams) (*model.Users, error) {
	if err := authValidator(params.Mail, params); err != nil {
		return nil, err
	}
	tx, err := s.repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	passwd := middleware.HashPassword(params.Password)
	entity := model.Users{
		Name:      params.Name,
		Mail:      params.Mail,
		Passwords: passwd,
		CreatedAt: time.Now().UTC(),
		UUID:      params.UUID,
	}
	if err := s.repo.Register(ctx, &entity); err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
