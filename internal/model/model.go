package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Status int

const (
	StatusPending Status = iota + 1
	StatusInProgress
	StatusDone
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending:
		return true
	case StatusInProgress:
		return true
	case StatusDone:
		return true
	}
	return false
}

type Hotels struct {
	UUID        uuid.UUID  `db:"uuid"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Services    pq.StringArray   `db:"services"`
	Prix        uint       `db:"prix"`
	Status      Status     `db:"status"`
	Ouverture   string     `db:"ouverture"`
	CreatedBy   string     `db:"created_by"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
