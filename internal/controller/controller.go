package controller

import (
	"newbier-hackglobal/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewController(app *fiber.App, db *gorm.DB) {
	u := usecase.NewUsecase(db)
	
	// ROUTE: Client
	app.Get("/",test(u))

	// ROUTE: Endpoint (IF NEED)
	// api := app.Group("/api")
	// api.Get("", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })	
}

func test(u *usecase.Usecase) fiber.Handler{
	return func(c *fiber.Ctx)error{
		data := u.GetUsecase()

		return c.Render("index", fiber.Map{
			"Title": data,
		}, "layouts/main")
	}
}