package app

import (
	"WebSocket/internal/endpoints"
	"WebSocket/internal/repository"
	"WebSocket/internal/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	app        *fiber.App
	endpoints  *endpoints.Endpoints
	services   *services.Services
	repository *repository.Repository
}

func New() *App {
	a := &App{}

	a.app = fiber.New()
	a.app.Use(cors.New(), logger.New(), recover.New())
	db := repository.CreateTable()

	a.repository = repository.New(db)
	a.services = services.New(a.repository)
	a.endpoints = endpoints.New(a.services)

	a.routers()
	return a
}

func (a *App) routers() {
	a.app.Post("/registration", a.endpoints.Registration)
}
func (a *App) Run() {

	a.app.Listen(os.Getenv("port"))

}
