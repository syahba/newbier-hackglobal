package controller

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
	"newbier-hackglobal/internal/usecase"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/database/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"fmt"
)

type Controller struct {
	usecase usecase.Usecase
}

func NewController(app *fiber.App, db *gorm.DB, ai *chatgpt.Model) {

	c := &Controller{
		usecase: *usecase.NewUsecase(db, ai),
	}

	// logger
	app.Use(logger.New())

	// ROUTE: Client
	app.Get("/", c.getUsecase)

	// ROUTE: API
	api := app.Group("/api")
	api.Get("/chat", c.getChat)
	api.Post("/chat", c.postChat)

	api.Get("/generate-itinerary", c.generateItinerary)
	api.Get("/destinations", c.getDestinations)
	api.Get("/itinerary/destinations", c.getItineraryDestination)
	api.Post("/itinerary/buddy", c.postItineraryBuddy)
	api.Put("/itinerary/buddy", c.putItineraryBuddy)
	api.Post("/itinerary/market", c.postItternaryMarket)
	api.Get("/itinerary/market/:id", c.getItternaryMarketByItternaryId)

	// 404 not found
	app.Use(func (c *fiber.Ctx)error{
		return c.Status(404).JSON(fiber.Map{
			"Message":"404 Not Found",
		})
	})
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

func (cn *Controller) generateItinerary(c *fiber.Ctx) error {

	data, _ := cn.usecase.GenerateItinerary()

	return c.Status(200).JSON(data)

}

func (cn *Controller) getItineraryDestination(c *fiber.Ctx) error {

	itineraryDestinationList, err := cn.usecase.GetItineraryDestination()

	if err != nil || len(itineraryDestinationList) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No itinerary destination Records",
		})
	}

	return c.Status(200).JSON(itineraryDestinationList)

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

	if err := c.BodyParser(&newChat); err != nil {
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
		"message": "successfully created new chat",
	})
}

func (cn *Controller) postItineraryBuddy(c *fiber.Ctx) error {
	var newItineraryBuddy model.ItineraryBuddy

	if err := c.BodyParser(&newItineraryBuddy); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	isItineraryIdZeroOrNegative := newItineraryBuddy.ItineraryID <= 0
	isUserIdZeroOrNegative := newItineraryBuddy.UserID <= 0
	isChatRoomIdZeroOrNegative := newItineraryBuddy.ChatRoomID <= 0
	isCreatedByEmpty := newItineraryBuddy.Description == "" || len(newItineraryBuddy.Description) == 0
	isDescriptionEmpty := newItineraryBuddy.Description == "" || len(newItineraryBuddy.Description) == 0

	if isItineraryIdZeroOrNegative || isUserIdZeroOrNegative || isCreatedByEmpty || isChatRoomIdZeroOrNegative || isDescriptionEmpty {
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id, user_id, chat_room_id, created_by, or description can't Zero, negative or Empty",
		})
	}

	if result, err := cn.usecase.GetItineraryById(newItineraryBuddy.ItineraryID); err != nil || result.ID == 0{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Itinerary with id %v was not found", newItineraryBuddy.ItineraryID),
		})
	}

	if result, err := cn.usecase.GetChatById(newItineraryBuddy.ChatRoomID); err != nil || result.ID == 0{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Chat Room with id %v was not found", newItineraryBuddy.ChatRoomID),
		})
	}

	if result, err := cn.usecase.GetUserById(newItineraryBuddy.UserID); err != nil || result.ID == 0{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("User with id %v was not found", newItineraryBuddy.UserID),
		})
	}

	err := cn.usecase.CreateItineraryBuddy(newItineraryBuddy)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Can't create itinerary buddy",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "successfully created new itinerary buddy",
	})
}

func (cn *Controller) putItineraryBuddy(c *fiber.Ctx) error{
	type customRequest struct{
		ItineraryId int `json:"itinerary_id"`
		UserId int 		`json:"user_id"`
	}

	var request customRequest

	if err := c.BodyParser(&request); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	if request.ItineraryId <= 0 || request.UserId <= 0{
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id or user_id can't zero or negative",
		})
	}

	if err := cn.usecase.JoinItineraryBuddy(request.UserId,request.ItineraryId); err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message": "can't join the itinerary buddy",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "successfully accept new itinerary buddy",
	})
}

func (cn *Controller) postItternaryMarket(c *fiber.Ctx) error {
	var newItternaryMarket model.ItineraryMarket

	if err := c.BodyParser(&newItternaryMarket); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	if newItternaryMarket.ItineraryID <= 0 || newItternaryMarket.DestinationProductID <= 0 || newItternaryMarket.Amount == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id, destination_product_id, or amount can't empty or negative",
		})
	}

	if result,err := cn.usecase.GetItineraryById(newItternaryMarket.ItineraryID);err != nil || result.ID == 0{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Itinerary with id %v was not found", newItternaryMarket.ItineraryID),
		})
	}

	if result,err := cn.usecase.GetDestinationProductById(newItternaryMarket.DestinationProductID);err != nil || result.ID == 0{
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Destination Product with id %v was not found", newItternaryMarket.DestinationProductID),
		})
	}

	if err := cn.usecase.CreateItternaryMarkets(newItternaryMarket); err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message": "Can't create Itternary Market",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "successfully created new itinerary market",
	})
}

func (cn *Controller) getItternaryMarketByItternaryId(c *fiber.Ctx) error {
	id,err := strconv.Atoi(c.Params("id"))

	if err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message": "path variable 'id' is invalid, make sure it is a INT",
		})
	}

	if id <= 0{
		return c.Status(400).JSON(fiber.Map{
			"message": "path variable 'id' can't zero or negative",
		})
	}

	ItternaryMarketList, err := cn.usecase.GetItternaryMarketByItternaryId(id)

	if err != nil || len(ItternaryMarketList) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No Itternary Market list Records",
		})
	}

	return c.Status(200).JSON(ItternaryMarketList)
}