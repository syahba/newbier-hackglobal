package main

import (
	"fmt"
	"newbier-hackglobal/cmd/generate/schema"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	model := chatgpt.GetModel(cfg.ChatGPTKey)
	result, _ := model.Generate(schema.AdditionalDestination("MINT Museum of Toys"))

	fmt.Println(result)
}
