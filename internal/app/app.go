package app

import (
	"WebSocket/internal/endpoints"
	"WebSocket/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	app       *fiber.App
	endpoints *endpoints.Endpoints
	services  *services.Services
}

func New() *App {
	a := &App{}

	a.app = fiber.New()
	a.app.Use(cors.New())
	a.services = services.New()
	a.endpoints = endpoints.New(a.services)

	a.routers()
	return a
}

func (a *App) routers() {
	a.app.Post("/registration", a.endpoints.Registration)
}
func (a *App) Run() {

	a.app.Listen(":8080")

}
