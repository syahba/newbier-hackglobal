package main

import (
	"fmt"
	"newbier-hackglobal/internal/controller"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Package
	// db, err := database.NewPostgresDB(cfg.DatabaseURL)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to load database: %v", err))
	// }
	// ai := chatgpt.GetModel(cfg.ChatGPTKey)

	// Server
	engine := jet.New("./views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	controller.NewController(app, &gorm.DB{}, &chatgpt.Model{})

	app.Listen(":" + cfg.Port)

}
