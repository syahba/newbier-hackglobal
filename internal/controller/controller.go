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
	app.Get("/destinations",getDestinations(u))
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

func getDestinations(u *usecase.Usecase) fiber.Handler{
	return func(c *fiber.Ctx)error{
		destinationList,err := u.GetDestinations()

		if err != nil || len(destinationList) == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message":"No destination Records",
			})
		}

		return c.Status(200).JSON(destinationList)
	}
}