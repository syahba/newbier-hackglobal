package chatgpt

import (
	"github.com/sashabaranov/go-openai"
	"context"
	"fmt"
)

type Model struct{
	Client *openai.Client
}

func GetModel(apiKey string) Model {
	return Model{
		Client : openai.NewClient(apiKey),
	}
}

func (model *Model) Generate(input string)(string,error){
	client := model.Client
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "",err
	}

	return resp.Choices[0].Message.Content,nil
}