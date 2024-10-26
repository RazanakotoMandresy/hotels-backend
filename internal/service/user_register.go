package service

import (
	"context"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/RazanakotoMandresy/hotels-backend/middleware"
)

func (s Service) Register(ctx context.Context, params RegisterParams) (*model.Users, error) {
	if err := authValidator(params.Mail, params); err != nil {
		return nil, err	
	}
	tx, err := s.repo.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	// passwd, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	passwd, err := middleware.Encrypt(params.Password)
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
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
