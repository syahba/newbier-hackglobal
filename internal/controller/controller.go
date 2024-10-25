package controller

import (
	"newbier-hackglobal/internal/usecase"
	chatgpt "newbier-hackglobal/pkg/chatGPT"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	usecase usecase.Usecase
}

func NewController(app *fiber.App, db *gorm.DB, ai *chatgpt.Model) {

	c := &Controller{
		usecase: *usecase.NewUsecase(db, ai),
	}

	// ROUTE: Client
	app.Get("/", c.getUsecase)
	app.Get("/home", c.getHome)
	app.Get("/preferrence", c.getPreferrence)
	app.Get("/bridging/1", c.getBridging)

	// ROUTE: API
	api := app.Group("/api")
	api.Get("/destinations", c.getDestinations)

}

// CLIENT
func (cn *Controller) getUsecase(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("index", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getHome(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("home", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getPreferrence(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("preferrence2", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getBridging(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("bridging", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

// API
func (cn *Controller) getDestinations(c *fiber.Ctx) error {

	destinationList, err := cn.usecase.GetDestinations()

	if err != nil || len(destinationList) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No destination Records",
		})
	}

	return c.Status(200).JSON(destinationList)

}
