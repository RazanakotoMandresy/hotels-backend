package service

import (
	"context"
	"fmt"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type RegisterParams struct {
	Name     string `valid:"required"`
	Mail     string `valid:"required"`
	Password string `valid:"required"`
	UUID     uuid.UUID
}

func (s Service) Register(ctx context.Context, params RegisterParams) (*model.Users, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return nil, err
	}
	if isMail := govalidator.IsEmail(params.Mail); !isMail {
		return nil, fmt.Errorf("%v is not an valid mail", params.Mail)
	}
	tx, err := s.repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	passwd, err := middleware.HashPassword(params.Password)
	if err != nil {
		return nil, err
	}
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

	return &entity, nil
}
