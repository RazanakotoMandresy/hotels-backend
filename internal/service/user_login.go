package service

import (
	"context"
	"errors"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
)

func (s Service) Login(ctx context.Context, params LoginParams) (*model.Users, error) {
	if err := authValidator(params.Mail, params); err != nil {
		return nil, err
	}
	tx, err := s.repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	users, err := s.repo.Login(ctx, params.Mail)
	if err != nil {
		return nil, err
	}
	decriptedPassword, err := middleware.Decrypt(users.Passwords)
	if err != nil {
		return nil, err
	}
	if decriptedPassword !=  params.Password{
		return nil, errors.New("bad passwords please retry")
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return users, nil
}
