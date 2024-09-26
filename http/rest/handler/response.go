package handler

import (
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/google/uuid"
)
type errorResponse struct {
	Err string `json:"err"`
}
type arrayHotels struct {
	Res []model.Hotels `json:"res"`
}
type responsesHotels struct {
	UUID        uuid.UUID  `json:"uuid"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      int        `json:"status"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at,omitempty"`
	Prix        uint       `json:"prix"`
	Ouverture   string     `json:"ouverture"`
	Services    []string   `json:"services"`
}
type responseUsers struct {
	Users     model.Users `json:"user"`
	ResString string      `json:"res"`
}
type responseString struct {
	Res string `json:"res"`
}
