package main

import (
	"fmt"
	"newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/config"
)

func main(){
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	model := chatgpt.GetModel(cfg.GeminiKey)
	result,_ := model.Generate("Hello gpt")

	fmt.Println(result)
}