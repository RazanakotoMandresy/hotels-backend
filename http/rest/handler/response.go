package handler

import (
	"github.com/RazanakotoMandresy/hotels-backend/internal/model"
)

type errorResponse struct {
	Err string `json:"err"`
}
type arrayHotels struct {
	Res []model.Hotels `json:"res"`
}

type responseUsers struct {
	Users     *model.Users `json:"user"`
	ResString string       `json:"res"`
}
type responseString struct {
	Res string `json:"res"`
}
