package usecase

import (
	"newbier-hackglobal/cmd/generate/schema"

	"github.com/sashabaranov/go-openai"
)

func generateItineraryMessage() []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu Siapa",
		},
		// {
		// 	Role:    openai.ChatMessageRoleSystem,
		// 	Content: schema.ParseSchema(schemaGenerateItinerary),
		// },
	}
}

var schemaGenerateItinerary = schema.Schema{}
