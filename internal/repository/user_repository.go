package repository

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (r Repository) Register(ctx context.Context, entity *model.Users) error {
	query := `INSERT INTO users (uuid , name , mail , list_hotels , passwords , created_at, updated_at, deleted_at)
	 VALUES (:uuid, :name, :mail, :list_hotels, :passwords , :created_at, :updated_at, :deleted_at)`
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

func (r Repository) Login(ctx context.Context, mail string) (*model.Users, error) {
	entity := new(model.Users)
	query := `SELECT * FROM users WHERE mail = $1 AND deleted_at IS NULL`
	err := r.Db.GetContext(ctx, entity, query, mail)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
func (r Repository) UpdateUser(ctx context.Context, entity *model.Users) error {
	query := `UPDATE users SET 
	name = :name,
	mail = :mail,
	list_hotels = :list_hotels ,
	updated_at = :updated_at WHERE uuid = :uuid;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return err
}
func (r Repository) FindUserByUUID(ctx context.Context, uuid string) (*model.Users, error) {
	entity := new(model.Users)
	query := `SELECT * FROM users WHERE uuid = $1 AND deleted_at IS NULL`
	if err := r.Db.GetContext(ctx, entity, query, uuid); err != nil {
		return nil, err
	}
	return entity, nil
}
