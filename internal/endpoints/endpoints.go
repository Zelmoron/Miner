package endpoints

import (
	"WebSocket/internal/requests"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Services interface {
	Registration(requests.UserRegRequest) error
	Login(requests.UserLoginRequest) error
}
type Endpoints struct {
	services Services
}

func New(services Services) *Endpoints {
	return &Endpoints{
		services: services,
	}
}

func (e *Endpoints) Registration(c *fiber.Ctx) error {

	var u requests.UserRegRequest
	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest - Bad Data",
		})
	}

	if err := validate.Struct(u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Validation error",
		})
	}

	err := e.services.Registration(u)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Registration failed",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "OK",
	})
}

func (e *Endpoints) Login(c *fiber.Ctx) error {
	var u requests.UserLoginRequest

	if err := c.BodyParser(&u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "BadRequest - Bad Data",
		})
	}

	if err := validate.Struct(u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Validation error",
		})
	}

	err := e.services.Login(u)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Login failed",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "OK",
	})

}
