package controller

import (
	"fmt"
	"newbier-hackglobal/internal/usecase"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/database/model"

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

	// ROUTE: API
	api := app.Group("/api")
	api.Get("/destinations", c.getDestinations)
	api.Get("/itternary/destinations",c.getItternaryDestination)

	api.Get("/chat",c.getChat)
	api.Post("/chat",c.postChat)

	api.Post("/itternary/buddy",c.postItternaryBuddy)
}

// CLIENT
func (cn *Controller) getUsecase(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("index", fiber.Map{
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

func (cn *Controller) getItternaryDestination(c *fiber.Ctx) error {

	itternaryDestinationList, err := cn.usecase.GetItternaryDestination()

	if err != nil || len(itternaryDestinationList) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No itternary destination Records",
		})
	}

	return c.Status(200).JSON(itternaryDestinationList)

}

func (cn *Controller) getChat(c *fiber.Ctx) error {

	chatList, err := cn.usecase.GetChat()

	if err != nil || len(chatList) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No chat list Records",
		})
	}

	return c.Status(200).JSON(chatList)
}

func (cn *Controller) postChat(c *fiber.Ctx) error {
	var newChat model.ChatRoom

	if err := c.BodyParser(&newChat); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	if newChat.Message == " " || len(newChat.Message) == 0 || newChat.CreatedBy == " " || len(newChat.CreatedBy) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Message or Created field can't empty",
		})
	}

	err := cn.usecase.CreateChat(newChat)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No destination Records",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":"successfully created new chat",
	})
}

func (cn *Controller) postItternaryBuddy(c *fiber.Ctx) error {
	var newItternaryBuddy model.ItineraryBuddy

	if err := c.BodyParser(&newItternaryBuddy); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	isItternaryIdZero := newItternaryBuddy.ItineraryID == 0
	isUserIdZero := newItternaryBuddy.UserID == 0
	isChatRoomId := newItternaryBuddy.ChatRoomID == 0
	isCreatedByEmpty := newItternaryBuddy.Description == "" || len(newItternaryBuddy.Description) == 0
	isDescriptionEmpty := newItternaryBuddy.Description == "" || len(newItternaryBuddy.Description) == 0

	if isUserIdZero || isChatRoomId || isCreatedByEmpty || isItternaryIdZero || isDescriptionEmpty {
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id, user_id, chat_room_id, created_by, or description is Zero or Empty",
		})
	}

	if _,err := cn.usecase.GetItternaryById(newItternaryBuddy.ItineraryID); err != nil{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Itternary with id %v was not found",newItternaryBuddy.ItineraryID),
		})
	}

	if _,err := cn.usecase.GetChatById(newItternaryBuddy.ChatRoomID); err != nil{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Chat Room with id %v was not found",newItternaryBuddy.ChatRoomID),
		})
	}

	if _,err := cn.usecase.GetUserById(newItternaryBuddy.UserID); err != nil{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("User with id %v was not found",newItternaryBuddy.UserID),
		})
	}

	err := cn.usecase.CreateItternaryBuddy(newItternaryBuddy)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Can't create itternary buddy",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":"successfully created new itternary buddy",
	})
}