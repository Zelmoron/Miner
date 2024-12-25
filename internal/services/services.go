package services

import (
	"WebSocket/internal/repository"
	"WebSocket/internal/requests"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
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

var jwtSecret = []byte("your_jwt_secret")         // Замените на свой секрет
var refreshSecret = []byte("your_refresh_secret") // Замените на свой секрет

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

	pass, err := s.database.GetUserLogin(user.Email)
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

	accessToken, err := generateJWT(user.Email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(user.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

func generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Функция для генерации refresh токена
func generateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // Срок действия 7 дней
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}
