package route

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"
	"github.com/gofiber/template/pug"
	"github.com/koungkub/wongnai/internal/handler"
	"github.com/koungkub/wongnai/internal/model"
)

func routing(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) { c.SendString("hello, world!") })
	app.Get("/reviews/:id", handler.GetReviewByID())
	app.Get("/reviews", handler.SearchReviewByQuery())
	app.Put("/reviews/:id", handler.EditReview())
}

// New get routing instance
func New(conf *model.Conf) *fiber.App {
	app := fiber.New()
	app.Settings.Views = pug.New("./public", ".pug")

	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(func(c *fiber.Ctx) {
		c.Locals("conf", conf)
		c.Next()
	})

	routing(app)

	return app
}
