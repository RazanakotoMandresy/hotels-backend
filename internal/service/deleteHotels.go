package service

import (
	"context"
	"errors"
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/middleware"
)

func (s Service) Delete(ctx context.Context, uuid string) error {
	hotels, err := s.GetHotel(ctx, uuid)
	if err != nil {
		return err
	}
	userUUID := middleware.GetUserUUIDInAuth(ctx)
	if userUUID == "" {
		return errors.New("no uuid in bearer auth")
	}
	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	if userUUID != hotels.CreatedBy {
		return errors.New("you are not the creator of this hotels")
	}
	now := time.Now().UTC()
	hotels.DeletedAt = &now
	err = s.repo.Update(ctx, *hotels)
	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return err
}
