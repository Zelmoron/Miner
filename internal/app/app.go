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
	"github.com/sirupsen/logrus"
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
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	a.app.Use(logger.New(), recover.New())

	public := a.app.Group("")
	public.Post("/registration", a.endpoints.Registration)
	public.Post("/login", a.endpoints.Login)
	public.Delete("/delete/:id", a.endpoints.Delete)

	jwt := a.app.Group("/jwt")
	jwt.Use(a.middleware.JWT)
	jwt.Get("/checktoken", a.endpoints.Check)

	refresh := a.app.Group("/refresh")
	refresh.Use(a.middleware.REFRESH)
	refresh.Get("/refresh", a.endpoints.Refresh)

}
func (a *App) Run() {
	logrus.Info("Start server")
	a.app.Listen(os.Getenv("port"))

}
