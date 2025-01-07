package services

import (
	"WebSocket/internal/repository"
	"WebSocket/internal/requests"
	"WebSocket/internal/utils"
	"errors"
	"strconv"

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

	err := s.database.GetUser(user.Email)
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

	//git
	// hash_veriry, err := password.Verify(hash, u)
	// if err != nil {
	// 	log.Print(err)
	// }

	return nil
}

func (s *Services) Login(user requests.UserLoginRequest) (string, string, error) {
	pass, id, err := s.database.GetUserLogin(user.Email)
	if err != nil {
		logrus.Info("User not found")
		return "", "", errors.New("Not found")
	}

	logrus.Info("Email found")

	hash_password, err := password.Verify(pass, user.Password)
	if err != nil {
		logrus.Error(err)
		return "", "", err
	}

	if !hash_password {
		logrus.Info(hash_password)
		return "", "", errors.New("Password false")
	}

	accessToken, err := utils.GenerateJWT(id)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	err = s.database.AddToken(refreshToken, id)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil

}

func (s *Services) NewJWT(id interface{}) (string, error) {
	idS := id.(string)
	ID, err := strconv.Atoi(idS)
	if err != nil {
		return "", err
	}
	accessToken, err := utils.GenerateJWT(ID)
	if err != nil {
		return "", err
	}
	return accessToken, nil

}
