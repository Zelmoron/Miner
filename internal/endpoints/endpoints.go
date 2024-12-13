package endpoints

import (
	"WebSocket/internal/requests"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Services interface {
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
			"status": "BadRequest",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "OK",
	})
}
