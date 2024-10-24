package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type LoginParams struct {
	Mail     string `valid:"required"`
	Password string `valid:"required"`
}
type CreateParams struct {
	Name        string `valid:"required"`
	Description string `valid:"required"`
	Services    pq.StringArray
	Prix        uint `valid:"required"`
	Status      bool
	Place       string `valid:"required"`
}
type RegisterParams struct {
	Name     string `valid:"required"`
	Mail     string `valid:"required"`
	Password string `valid:"required"`
	UUID     uuid.UUID
}

type UpdateParams struct {
	UUID             string `valid:"required"`
	Name             *string
	Description      *string
	Prix             *uint
	Status           *bool
	ReservationLists *[]string
	UpdatedAt        time.Time
}
type FilterParams struct {
	Name      string
	Ouverture string
	Place     string
	Service   []string
	MinBudget uint
	MaxBudget uint
}
type ReserveParams struct {
	Starting_date string
	Ending_date   string
	Password      string
}
