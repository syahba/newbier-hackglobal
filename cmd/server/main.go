package main

import (
	"fmt"
	"newbier-hackglobal/internal/controller"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Package
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load database: %v", err))
	}
	ai := chatgpt.GetModel(cfg.ChatGPTKey)

	// Server
	engine := jet.New("./views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New())
	app.Static("/public", "./public")
	controller.NewController(app, db, ai)

	app.Listen("localhost:" + cfg.Port)

}
