package services

import (
	"WebSocket/internal/repository"
	"WebSocket/internal/requests"
	"errors"

	"github.com/sirupsen/logrus"
	password "github.com/vzglad-smerti/password_hash"
)

type Services struct {
	database *repository.Repository
}

func New(database *repository.Repository) *Services {
	return &Services{
		database: database,
	}
}

func (s *Services) Registration(user requests.UserRegRequest) error {

	err := s.database.GetUser(user)
	if err == nil {
		logrus.Info("User exists")
		return errors.New("Already exists(user)")
	}

	//hash password
	hash, err := password.Hash(user.Password)
	if err != nil {
		logrus.Println(err)
	}

	err = s.database.CreateUser(user, hash)

	if err != nil {
		logrus.Info("Insert failed")
		return err
	}

	// hash_veriry, err := password.Verify(hash, u)
	// if err != nil {
	// 	log.Print(err)
	// }

	return nil
}

func JWT() {

}

func Refresh() {

}
