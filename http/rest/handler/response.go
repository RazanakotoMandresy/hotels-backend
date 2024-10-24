package handler

import (
	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

type errorResponse struct {
	Err string `json:"err"`
}
type arrayHotelsResponse struct {
	Res []model.Hotels `json:"res"`
}

type responseUsers struct {
	Users *model.Users `json:"user"`
	// Resstring most of the time will be the token
	ResString string `json:"res"`
}
type responseString struct {
	Res string `json:"res"`
}
type responseHotel struct {
	Res model.Hotels `json:"res"`
}
