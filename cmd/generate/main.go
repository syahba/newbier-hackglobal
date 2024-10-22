package main

import (
	"fmt"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/gemini"
)

func main(){
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	gemini := gemini.GetModel(cfg.GeminiKey)

	result,_ := gemini.Chat("give me the recipes of pizza")

	fmt.Println(result)
}