package endpoints

import (
	"WebSocket/internal/requests"
	"net/http"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type Services interface {
	Registration(requests.UserRegRequest) error
	Login(requests.UserLoginRequest) (string, string, error)
	NewJWT(interface{}) (string, error)
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

	access_token, refresh_token, err := e.services.Login(u)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Login failed",
		})
	}
	ch := make(chan bool)

	// Установка токенов в HTTP-only куки - так потому что на разныж хостах
	go func(с *fiber.Ctx, ch chan bool) {
		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    access_token,
			Expires:  time.Now().Add(time.Hour * 100 * 100),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "None",
		})

		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    refresh_token,
			Expires:  time.Now().Add(time.Hour * 100 * 100),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "None",
		})

		ch <- true
	}(c, ch)

	select {
	case <-ch:

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": "OK", "access_token": access_token,
		})
	case <-time.After(2 * time.Second):

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "LongCookie",
		})
	}

}

func (e *Endpoints) Check(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"name": c.Locals("sub"),
	})
}

func (e *Endpoints) Refresh(c *fiber.Ctx) error {
	id := c.Locals("sub")
	access, err := e.services.NewJWT(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"refresh": "Bad",
		})
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(c *fiber.Ctx, access string, wg *sync.WaitGroup) {
		defer wg.Done()
		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    access,
			Expires:  time.Now().Add(time.Hour * 100 * 100),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "None",
		})
	}(c, access, &wg)

	wg.Wait()

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"name": c.Locals("sub"),
	})
}
