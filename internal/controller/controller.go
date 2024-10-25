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
	app.Get("/home", c.getHome)
	app.Get("/preference/1", c.getPreference)
	app.Get("/preference/2", c.getPreference2)
	app.Get("/preference/3", c.getPreference3)
	app.Get("/bridging/1", c.getBridging)
	app.Get("/bridging/2", c.getBridging2)
	app.Get("/bridging/3", c.getBridging3)
	app.Get("/itinerary/1", c.getItinerary)
	app.Get("/itinerary/2", c.getItinerary2)
	app.Get("/market", c.getMarket)
	app.Get("/transaction", c.getTransaction)
	app.Get("/destination", c.getDestination)
	app.Get("/chat", c.getChat)
	app.Get("/match", c.getMatch)
	app.Get("/buddy", c.getBuddy)
	app.Get("/buddy/profile", c.getBuddyProfile)

	// ROUTE: API
	api := app.Group("/api")
	api.Get("/generate-itinerary", c.generateItinerary)

	api.Get("/destinations", c.getDestinations)
	api.Get("/itinerary/destinations", c.getItineraryDestination)

	api.Get("/chat", c.getChat)
	api.Post("/chat", c.postChat)

	api.Post("/itinerary/buddy", c.postItineraryBuddy)
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

func (cn *Controller) getPreference(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("preference", fiber.Map{
		"Title":  data,
		"Action": "Confirm",
	}, "layouts/main")
}

func (cn *Controller) getPreference2(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("preference-2", fiber.Map{
		"Title":  data,
		"Action": "Confirm",
	}, "layouts/main")
}

func (cn *Controller) getPreference3(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("preference-3", fiber.Map{
		"Title":  data,
		"Action": "Confirm",
	}, "layouts/scroll")
}

func (cn *Controller) getBridging(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("bridging", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getBridging2(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("bridging-2", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getBridging3(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("bridging-3", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getItinerary(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("itinerary", fiber.Map{
		"Title":  data,
		"Action": "Next",
	}, "layouts/scroll")
}

func (cn *Controller) getItinerary2(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("home-itinerary", fiber.Map{
		"Title":  data,
		"Action": "Completed",
	}, "layouts/scroll")
}

func (cn *Controller) getMarket(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("market", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getTransaction(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("transaction", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getDestination(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("destination", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getChat(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("chat-room", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getMatch(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("match", fiber.Map{
		"Title": data,
	}, "layouts/scroll")
}

func (cn *Controller) getBuddy(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("buddy", fiber.Map{
		"Title": data,
	}, "layouts/main")
}

func (cn *Controller) getBuddyProfile(c *fiber.Ctx) error {

	data := cn.usecase.GetUsecase()

	return c.Render("buddy-profile", fiber.Map{
		"Title":  data,
		"Action": "Confirm",
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

	isItineraryIdZero := newItineraryBuddy.ItineraryID == 0
	isUserIdZero := newItineraryBuddy.UserID == 0
	isChatRoomId := newItineraryBuddy.ChatRoomID == 0
	isCreatedByEmpty := newItineraryBuddy.Description == "" || len(newItineraryBuddy.Description) == 0
	isDescriptionEmpty := newItineraryBuddy.Description == "" || len(newItineraryBuddy.Description) == 0

	if isUserIdZero || isChatRoomId || isCreatedByEmpty || isItineraryIdZero || isDescriptionEmpty {
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id, user_id, chat_room_id, created_by, or description is Zero or Empty",
		})
	}

	if _, err := cn.usecase.GetItineraryById(newItineraryBuddy.ItineraryID); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Itinerary with id %v was not found", newItineraryBuddy.ItineraryID),
		})
	}

	if _, err := cn.usecase.GetChatById(newItineraryBuddy.ChatRoomID); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Chat Room with id %v was not found", newItineraryBuddy.ChatRoomID),
		})
	}

	if _, err := cn.usecase.GetUserById(newItineraryBuddy.UserID); err != nil {
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
