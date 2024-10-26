package repository

import (
	"context"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

func (r Repository) CreateReservation(ctx context.Context, entity *model.Reservation) error {
	query := `INSERT INTO reservation (uuid , reserved_by_uuid , hotels_uuid , reservation_date_start , reservation_date_end )  
	VALUES (:uuid , :reserved_by_uuid , :hotels_uuid , :reservation_date_start , :reservation_date_end )`
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
	return nil
}
