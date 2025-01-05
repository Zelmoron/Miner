package middleware

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

var jwtSecret = []byte("your_jwt_secret")
var refreshSecret = []byte("your_refresh_secret")

func (m *Middleware) JWT(c *fiber.Ctx) error {

	tokenString := c.Cookies("access_token")
	if tokenString == "" {
		logrus.Warn("Token is missing")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token is missing",
		})
	}
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм токена - HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Warn("invalid signing method")
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorMalformed)
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		logrus.Warn("Invalid token - time or error with parse")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Извлекаем sub из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logrus.Warn("Invalid token claims")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		logrus.Warn("Missing sub in token")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing sub in token",
		})
	}

	// Передаем sub в контекст
	c.Locals("sub", sub)

	logrus.Info("JWT middleware passed, sub: ", sub)
	// fmt.Println(c.Cookies("access_token"), 1)

	return c.Next()

}

func (m *Middleware) REFRESH(c *fiber.Ctx) error {

	tokenString := c.Cookies("refresh_token")
	fmt.Println(tokenString)
	if tokenString == "" {
		logrus.Warn("Token is missing")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token is missing",
		})
	}

	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм токена - HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Warn("invalid signing method")
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorMalformed)
		}
		return refreshSecret, nil
	})

	if err != nil || !token.Valid {
		logrus.Warn("Invalid token - time or error with parse")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Извлекаем sub из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logrus.Warn("Invalid token claims")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		logrus.Warn("Missing sub in token")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing sub in token",
		})
	}

	// Передаем sub в контекст
	c.Locals("sub", sub)

	logrus.Info("Refresh middleware passed, sub: ", sub)
	// fmt.Println(c.Cookies("access_token"), 1)

	return c.Next()

}
