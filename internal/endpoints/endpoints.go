package endpoints

import "github.com/gofiber/fiber/v2"

type Endpoints struct {
}

func New() *Endpoints {
	return &Endpoints{}
}

func (e *Endpoints) Registration(c *fiber.Ctx) error {

}
