package controller

import (
	"fmt"
	internal_model "newbier-hackglobal/internal/model"
	"newbier-hackglobal/internal/usecase"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/database/model"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
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
	api.Get("/chat", c.getChat)
	api.Post("/chat", c.postChat)
	api.Get("/itinerary", c.getItineraries)
	api.Post("/general-matter",c.generalMatter)
	api.Get("/generate-itinerary/destination", c.generateItineraryWithDestination)
	api.Get("/generate-itinerary", c.generateItinerary)
	api.Get("/destinations", c.getDestinations)
	api.Get("/destinations/:id", c.getDestinationById)
	api.Get("/itinerary/destinations", c.getItineraryDestination)
	api.Post("/itinerary/buddy", c.postItineraryBuddy)
	api.Put("/itinerary/buddy", c.putItineraryBuddy)
	api.Post("/itinerary/market", c.postItternaryMarket)
	api.Get("/itinerary/market/:id", c.getItternaryMarketByItternaryId)
	// 404 not found
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"Message": "404 Not Found",
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

func (cn *Controller) generalMatter(c *fiber.Ctx) error{
	var body internal_model.Combo

	if err := c.BodyParser(&body); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message":"Can't parse body",
		})
	}
	
	//fmt.Println(body)

	if len(body.Trip) == 0 || body.Trip == " " || body.User_id <= 0|| len(body.Itinerary) == 0{
		//fmt.Println(len(body.Trip) == 0, body.Trip == " " , body.User_id <= 0 , len(body.Itinerary) == 0)
		return c.Status(400).JSON(fiber.Map{
			"message":"field trip, userId or itinerary can't empty, zero or negative",
		})
	}

	body.IsBuddy = true
	if len(body.Destination) == 0{
		body.Destination = "empty"
	}

	if len(body.Activity) == 0 {
		body.Activity = "empty"
	}

	if len(body.Description) == 0 {
		body.Description = "empty"
	}

	if err := cn.usecase.CreateItineraryAndItineraryDestinationAndBuddy(body); err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message":"something went wrong, cant create itinerary, itineraryDestination and itineraryBuddy",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":"successfully create new itinerary, itineraryDestination and itineraryBuddy",
	})
}

func (cn *Controller) getItineraries(c *fiber.Ctx) error {

	itineraryList, err := cn.usecase.GetItinerary()
	if err != nil || len(itineraryList) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No Itinerary Records",
		})
	}
	var itineraryResponseList []internal_model.ItineraryResponse

	for _, itinerary := range itineraryList {
		itineraryResponse := internal_model.ItineraryResponse{
			ItineraryID:   int(itinerary.ID),
			ItineraryTime: groupDestinationsByTime(itinerary.ItineraryDestinations),
		}
		itineraryResponseList = append(itineraryResponseList, itineraryResponse)
	}

	return c.Status(200).JSON(itineraryResponseList)
}

func groupDestinationsByTime(destinations []model.ItineraryDestination) []internal_model.ItineraryTime {
	timeGroups := map[string][]model.Destination{
		"morning":   {},
		"afternoon": {},
		"night":     {},
	}

	for _, dest := range destinations {
		timeOfDay := getTimeOfDay(dest.Time)
		timeGroups[timeOfDay] = append(timeGroups[timeOfDay], dest.Destination)
	}

	var itineraryTimes []internal_model.ItineraryTime
	for time, dests := range timeGroups {
		if len(dests) > 0 {
			itineraryTimes = append(itineraryTimes, internal_model.ItineraryTime{
				Time:         time,
				Destinations: dests,
			})
		}
	}

	return itineraryTimes
}

func getTimeOfDay(timeStr string) string {

	if strings.Contains(timeStr, "AM") {
		hour := strings.Split(timeStr, ":")[0]
		if hour == "07" || hour == "08" || hour == "09" || hour == "10" {
			return "morning"
		}
		return "night"
	} else if strings.Contains(timeStr, "PM") {
		hour := strings.Split(timeStr, ":")[0]
		if hour == "12" || hour == "01" || hour == "02" {
			return "afternoon"
		}
		return "night"
	}
	return "night"
}

func (cn *Controller) getDestinations(c *fiber.Ctx) error {
	search := c.Query("search")

	destinationList, _ := cn.usecase.GetDestinations(search)

	return c.Status(200).JSON(destinationList)

}

func (cn *Controller) getDestinationById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "params id must number",
		})
	}

	data, err := cn.usecase.GetDestinationById(id)

	if err != nil || data.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Destination with id %v was not found", id),
		})
	}

	return c.Status(200).JSON(data)
}

func (cn *Controller) generateItinerary(c *fiber.Ctx) error {
	activity := c.Query("activity")
	trip := c.Query("trip")

	if activity == "_" || len(activity) == 0 || trip == "_" || len(trip) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "query with field 'activity' and 'trip' can't empty",
		})
	}

	data, _ := cn.usecase.GenerateItinerary(activity, trip)

	return c.Status(200).JSON(data)
}

func (cn *Controller) generateItineraryWithDestination(c *fiber.Ctx) error {
	destination := c.Query("destination")
	trip := c.Query("trip")

	if destination == "_" || len(destination) == 0 || trip == "_" || len(trip) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "query with field 'destination' and 'trip' can't empty",
		})
	}

	data, _ := cn.usecase.GenerateItineraryWithDestination(destination, trip)

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

func (cn *Controller) getChats(c *fiber.Ctx) error {

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

	if result, err := cn.usecase.GetItineraryById(newItineraryBuddy.ItineraryID); err != nil || result.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Itinerary with id %v was not found", newItineraryBuddy.ItineraryID),
		})
	}

	if result, err := cn.usecase.GetChatById(newItineraryBuddy.ChatRoomID); err != nil || result.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Chat Room with id %v was not found", newItineraryBuddy.ChatRoomID),
		})
	}

	if result, err := cn.usecase.GetUserById(newItineraryBuddy.UserID); err != nil || result.ID == 0 {
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

func (cn *Controller) putItineraryBuddy(c *fiber.Ctx) error {
	type customRequest struct {
		ItineraryId int `json:"itinerary_id"`
		UserId      int `json:"user_id"`
	}

	var request customRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	if request.ItineraryId <= 0 || request.UserId <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id or user_id can't zero or negative",
		})
	}

	if err := cn.usecase.JoinItineraryBuddy(request.UserId, request.ItineraryId); err != nil {
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

	if err := c.BodyParser(&newItternaryMarket); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Can't parse the body",
		})
	}

	if newItternaryMarket.ItineraryID <= 0 || newItternaryMarket.DestinationProductID <= 0 || newItternaryMarket.Amount == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "field itinerary_id, destination_product_id, or amount can't empty or negative",
		})
	}

	if result, err := cn.usecase.GetItineraryById(newItternaryMarket.ItineraryID); err != nil || result.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Itinerary with id %v was not found", newItternaryMarket.ItineraryID),
		})
	}

	if result, err := cn.usecase.GetDestinationProductById(newItternaryMarket.DestinationProductID); err != nil || result.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("Destination Product with id %v was not found", newItternaryMarket.DestinationProductID),
		})
	}

	if err := cn.usecase.CreateItternaryMarkets(newItternaryMarket); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Can't create Itternary Market",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "successfully created new itinerary market",
	})
}

func (cn *Controller) getItternaryMarketByItternaryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "path variable 'id' is invalid, make sure it is a INT",
		})
	}

	if id <= 0 {
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
