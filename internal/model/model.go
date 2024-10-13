package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Hotels struct {
	UUID        uuid.UUID      `db:"uuid"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
	Services    pq.StringArray `db:"services"`
	Prix        uint           `db:"prix"`
	Status      bool           `db:"status"`
	Place       string         `db:"place"`
	Ouverture   string         `db:"ouverture"`
	CreatedBy   string         `db:"created_by"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   *time.Time     `db:"updated_at"`
	DeletedAt   *time.Time     `db:"deleted_at"`
	Images      pq.StringArray `db:"images"`
}
type Users struct {
	UUID       uuid.UUID      `db:"uuid"`
	Name       string         `db:"name"`
	Passwords  string         `db:"passwords"`
	Mail       string         `db:"mail"`
	ListHotels pq.StringArray `db:"list_hotels"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  *time.Time     `db:"updated_at"`
	DeletedAt  *time.Time     `db:"deleted_at"`
}
