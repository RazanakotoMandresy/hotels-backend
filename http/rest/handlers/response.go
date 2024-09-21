package handlers

import (
	"time"

	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
	"github.com/google/uuid"
)

type errorResponse struct {
	Err string `json:"err"`
}
type response struct {
	UUID        uuid.UUID    `json:"uuid"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      model.Status `json:"status"`
	CreatedOn   time.Time    `json:"created_on"`
	UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
}
type modelResponse struct {
	Res model.Hotels `json:"res"`
}
type responseString struct {
	Res string `json:"res"`
}
