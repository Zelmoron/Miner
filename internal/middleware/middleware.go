package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

var jwtSecret = []byte("your_jwt_secret")

func (m *Middleware) JWT(c *fiber.Ctx) error {

	// tokenString := c.Cookies("access_token")
	// if tokenString == "" {
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "Token is missing",
	// 	})
	// }
	// fmt.Println(tokenString)
	// // Парсим токен
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Проверяем, что алгоритм токена - HMAC
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorMalformed)
	// 	}
	// 	return jwtSecret, nil
	// })

	// if err != nil || !token.Valid {
	// 	logrus.Warn("Invalid token")
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Invalid token",
	// 	})
	// }

	// // Извлекаем sub из токена
	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok || !token.Valid {
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Invalid token claims",
	// 	})
	// }

	// sub, ok := claims["sub"].(string)
	// if !ok {
	// 	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Missing sub in token",
	// 	})
	// }

	// // Передаем sub в контекст
	// c.Locals("sub", sub)

	// logrus.Info("JWT middleware passed, sub: ", sub)
	fmt.Println(c.Cookies("john"), 1)

	return c.Next()

}
