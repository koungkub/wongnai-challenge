package route

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"
)

// New get routing instance
func New() *fiber.App {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{fiber.MethodGet, fiber.MethodPost},
		AllowHeaders:     []string{fiber.HeaderAuthorization, fiber.HeaderContentType},
		ExposeHeaders:    []string{fiber.HeaderAuthorization, fiber.HeaderContentType},
		AllowCredentials: false,
		MaxAge:           60 * 60,
	}))

	return app
}
