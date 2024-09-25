package repository

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (r Repository) Register(ctx context.Context, entity *model.Users) error {
	query := `INSERT INTO users (uuid , name , mail , list_hotels , created_at, updated_at, deleted_at)
	 VALUES (:uuid, :name, :mail, :list_hotels, :created_at, :updated_at, :deleted_at)`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return err
		}
	}
	return err
}

func (r Repository) Login(ctx context.Context, mail, passwords, uuid string) (*model.Users, error) {
	entity := new(model.Users)
	query := `SELECT * FROM users WHERE mail = $1  AND passwords = $2 AND uuid = $3 AND deleted_at IS NULL`
	err := r.Db.GetContext(ctx, entity, query, mail, passwords, uuid)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
