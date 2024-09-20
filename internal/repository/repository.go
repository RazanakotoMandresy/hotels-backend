package repository

import (
	"context"

	"github.com/RazanakotoMandresy/deliveryapp-backend/internal/model"
	"github.com/RazanakotoMandresy/deliveryapp-backend/pkg/db"
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
	return entity, db.HandleError(err)
}

func (r Repository) Create(ctx context.Context, entity *model.Hotels) error {
	query := `INSERT INTO hotels (uuid ,name, description, status, created_on, updated_on)
                VALUES (:uuid ,:name, :description, :status, :created_on, :updated_on) RETURNING uuid;`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return db.HandleError(err)
		}
	}
	return db.HandleError(err)
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
	return db.HandleError(err)
}

func (r Repository) FindAll(ctx context.Context) ([]model.Hotels, error) {
	var entities []model.Hotels
	query := "SELECT * FROM todo WHERE deleted_on IS NULL"
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, db.HandleError(err)
}
