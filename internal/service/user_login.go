package service

import (
	"context"

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

	users, err := s.repo.GetUser(ctx, params.Mail)
	if err != nil {
		return nil, err
	}

	if err := middleware.Decrypt(users.Passwords, params.Password); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return users, nil
}
