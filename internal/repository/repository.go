package repository

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}
func (r Repository) Find(ctx context.Context, uuid string) (model.Hotels, error) {
	entity := model.Hotels{}
	query := "SELECT * FROM hotels WHERE uuid = $1 AND deleted_on IS NULL"
	err := r.Db.GetContext(ctx, &entity, query, uuid)
	return entity, err
}

func (r Repository) Create(ctx context.Context, entity *model.Hotels) error {
	query := `INSERT INTO hotels (uuid ,name, description, services, prix, created_at, updated_at)
                VALUES (:uuid ,:name, :description, :services, :prix,:created_at, :updated_at) RETURNING uuid;`
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

func (r Repository) Update(ctx context.Context, entity model.Hotels) error {
	query := `UPDATE hotels
                SET name = :name, 
                    description = :description, 
                    status = :status, 
                    created_on = :created_on, 
                    updated_on = :updated_on, 
                    deleted_on = :deleted_on
                WHERE uuid = :uuid;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return err
}

func (r Repository) FindAll(ctx context.Context) ([]model.Hotels, error) {
	var entities []model.Hotels
	query := "SELECT * FROM hotels WHERE deleted_on IS NULL"
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, err
}
