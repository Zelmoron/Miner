package endpoints

import (
	"WebSocket/internal/requests"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Services interface {
	Registration(requests.UserRegRequest)
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

	e.services.Registration(u)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "OK",
	})
}
