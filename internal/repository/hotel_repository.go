package repository

import (
	"context"
	"fmt"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (r Repository) Find(ctx context.Context, uuid string) (*model.Hotels, error) {
	entity := new(model.Hotels)
	query := "SELECT * FROM hotels WHERE uuid = $1 AND deleted_at IS NULL"
	err := r.Db.GetContext(ctx, entity, query, uuid)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// find only by uuid but without even it was deleted
func (r Repository) FindHotelsByUUID(ctx context.Context, uuid string) (*model.Hotels, error) {
	entity := new(model.Hotels)
	query := "SELECT * FROM hotels WHERE uuid = $1"
	err := r.Db.GetContext(ctx, entity, query, uuid)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r Repository) Create(ctx context.Context, entity *model.Hotels) error {
	query := `INSERT INTO hotels (uuid ,name, description, services, prix, ouverture, place,status,created_by ,  created_at, updated_at)
                VALUES (:uuid ,:name, :description, :services, :prix, :ouverture, :place,:status,:created_by,:created_at, :updated_at) RETURNING uuid;`
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
                    created_at = :created_at, 
                    updated_at = :updated_at, 
                    deleted_at = :deleted_at,
					images = :images
                WHERE uuid = :uuid;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return err
}

func (r Repository) FindAll(ctx context.Context) ([]model.Hotels, error) {
	var entities []model.Hotels
	query := "SELECT * FROM hotels WHERE deleted_at IS NULL"
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, err
}
func (r Repository) SearchQuery(ctx context.Context, search string) ([]model.Hotels, error) {
	entity := new([]model.Hotels)
	pattern := fmt.Sprint("%" + search + "%")
	err := r.Db.SelectContext(ctx, entity, "SELECT * FROM hotels WHERE name LIKE $1", pattern)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return *entity, err
}
func (r Repository) Filter(ctx context.Context, condFilter, column string) ([]model.Hotels, error) {
	entity := new([]model.Hotels)
	matche := fmt.Sprint("%" + condFilter + "%")
	query := fmt.Sprintf(`SELECT * FROM hotels WHERE %v LIKE $1`, column)
	err := r.Db.SelectContext(ctx, entity, query, matche)
	if err != nil {
		return nil, err
	}
	return *entity, nil
}
func (r Repository) FilterPrice(ctx context.Context, minBudget, maxBudget uint) ([]model.Hotels, error) {
	entity := new([]model.Hotels)
	query := fmt.Sprintf(`SELECT * FROM hotels WHERE prix >= %v AND prix <= %v AND deleted_at IS NULL `, minBudget, maxBudget)
	if err := r.Db.SelectContext(ctx, entity, query); err != nil {
		return nil, err
	}
	return *entity, nil
}
