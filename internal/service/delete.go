package service

import (
	"context"
	"time"
)

func (s Service) Delete(ctx context.Context, uuid string) error {
	hotels, err := s.Get(ctx, uuid)
	if err != nil {
		return err
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

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
