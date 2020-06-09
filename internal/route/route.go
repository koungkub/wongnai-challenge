package route

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"
	"github.com/gofiber/template/pug"
)

func routing(app *fiber.App) {

}

// New get routing instance
func New() *fiber.App {
	app := fiber.New()
	app.Settings.Templates = pug.New("./public", ".pug")

	app.Use(requestid.New())
	app.Use(recover.New())

	routing(app)

	return app
}
