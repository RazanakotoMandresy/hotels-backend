package service

import "github.com/RazanakotoMandresy/deliveryapp-backend/internal/repository"
type Service struct {
    repo repository.Repository
}

func NewService(r repository.Repository) Service {
    return Service{
        repo: r,
    }
}
