package controller

import (
	"newbier-hackglobal/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewController(app *fiber.App, db *gorm.DB) {
	u := usecase.NewUsecase(db)

	// ROUTE: Client
	app.Get("", func(c *fiber.Ctx) error {

		data := u.GetUsecase()

		return c.Render("index", fiber.Map{
			"Title": data,
		}, "layouts/main")
	})

	// ROUTE: Endpoint (IF NEED)
	// api := app.Group("/api")
	// api.Get("", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })
}
