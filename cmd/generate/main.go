package main

import (
	"fmt"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/config"

	"github.com/sashabaranov/go-openai"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	model := chatgpt.GetModel(cfg.ChatGPTKey)
	result, _ := model.Generate([]openai.ChatCompletionMessage{})

	fmt.Println(result)
}
