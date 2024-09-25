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

type modelResponse struct {
	Hotels model.Hotels `json:"resHotels"`
	Users  model.Users  `json:"resUsers"`
}
type responseString struct {
	Res string `json:"res"`
}
