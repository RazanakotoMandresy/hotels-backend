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
	Ouverture   string `valid:"required"`
	Place       string `valid:"required"`
}
type RegisterParams struct {
	Name     string `valid:"required"`
	Mail     string `valid:"required"`
	Password string `valid:"required"`
	UUID     uuid.UUID
}

type UpdateParams struct {
	UUID        string `valid:"required"`
	Name        *string
	Description *string
	Prix        *uint
	Status      *bool
	Ouverture   *string
	UpdatedAt   time.Time
}
type FilterParams struct {
	UUID      string   `valid:"required"`
	Name      string   `json:"name"`
	Ouverture string   `json:"ouverture"`
	Place     string   `json:"place"`
	Service   []string `json:"service"`
	Prix      uint     `json:"prix"`
}
