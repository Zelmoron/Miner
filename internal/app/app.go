package app

import (
	"WebSocket/internal/endpoints"
	"WebSocket/internal/services"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	app       *fiber.App
	endpoints *endpoints.Endpoints
	services  *services.Services
}

func New() *App {
	a := &App{}

	a.app = fiber.New()
	a.services = services.New()
	a.endpoints = endpoints.New()

	return a
}

func (a *App) routers() {
	a.app.Post("/registration")
}
func (a *App) Run() {

	a.app.Listen(":8080")

}
