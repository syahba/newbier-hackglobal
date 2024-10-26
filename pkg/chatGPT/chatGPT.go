package chatgpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Model struct {
	client *openai.Client
}

func GetModel(apiKey string) *Model {
	return &Model{
		client: openai.NewClient(apiKey),
	}
}

func (model *Model) Generate(messages []openai.ChatCompletionMessage) (string, error) {
	client := model.client
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT4oMini,
			Temperature: 0.8,
			Messages:    messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (model *Model) GenerateCustom(temperature float32, token int, messages []openai.ChatCompletionMessage) (string, error) {
	client := model.client
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT4oMini,
			Temperature: temperature,
			Messages:    messages,
			MaxTokens:   token,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
