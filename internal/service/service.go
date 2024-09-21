package service

import "github.com/RazanakotoMandresy/hotels-backend/internal/repository"
type    Service struct {
    repo repository.Repository
}

func NewService(r repository.Repository) Service {
    return Service{
        repo: r,
    }
}
