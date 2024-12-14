package services

import (
	"WebSocket/internal/repository"
	"WebSocket/internal/requests"
)

type Services struct {
	database *repository.Repository
}

func New(database *repository.Repository) *Services {
	return &Services{
		database: database,
	}
}

func (s *Services) Registration(user *requests.UserRegRequest) {

}
