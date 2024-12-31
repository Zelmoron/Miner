package app

import (
	"WebSocket/internal/endpoints"
	"WebSocket/internal/middleware"
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
	middleware *middleware.Middleware
}

func New() *App {
	a := &App{}

	a.app = fiber.New()

	db := repository.CreateTable()

	a.repository = repository.New(db)
	a.services = services.New(a.repository)
	a.endpoints = endpoints.New(a.services)

	a.middleware = middleware.New()

	a.routers()
	return a
}

func (a *App) routers() {

	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500", // Change this to your frontend's URL
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	a.app.Use(logger.New(), recover.New())

	public := a.app.Group("/")
	public.Post("/registration", a.endpoints.Registration)
	public.Post("/login", a.endpoints.Login)

	private := a.app.Group("/")
	private.Use(a.middleware.JWT)
	private.Get("/checktoken", a.endpoints.Check)
}
func (a *App) Run() {

	a.app.Listen(os.Getenv("port"))

}
