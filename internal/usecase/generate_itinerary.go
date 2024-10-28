package usecase

import (
	"fmt"
	"newbier-hackglobal/cmd/generate/schema"
	"github.com/sashabaranov/go-openai"
)

func GenerateItinerary(destinations []map[string]any, activity string, trip string) []openai.ChatCompletionMessage {
	data1 := destinations[:len(destinations)/2]
	data2 := destinations[len(destinations)/2:]
	var destinationData = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Berikut data destinasti-nya " + fmt.Sprint(data1),
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Berikut lanjutan data destinasti " + fmt.Sprint(data2),
		},
	}

	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu adalah assistant perekomendasi traveling itinerary destination",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu akan memilih destinasti mana saja yang sesuai dengan user inginkan dari pagi hingga malam",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `format waktunya akan seperti ini ["morning", "afternoon", "evening", "night"]`,
		},
		destinationData[0],
		destinationData[1],
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "destinasti yang saya berikan ini hanya destinasti dari negara singapura",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "pilihlah 2 hingga 4 destinasti id untuk setiap waktu-nya",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Saya ingin kegiatan %s dengan trip %s", activity, trip),
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "hanya berikan waktu dan destination id nya saja",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: schema.ParseSchema(schemaGenerateItinerary),
		},
	}
}

func GenerateItineraryWithDestination(destinations []map[string]any, destination string, trip string) []openai.ChatCompletionMessage {
	data1 := destinations[:len(destinations)/2]
	data2 := destinations[len(destinations)/2:]
	var destinationData = []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Berikut data destinasti-nya " + fmt.Sprint(data1),
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Berikut lanjutan data destinasti " + fmt.Sprint(data2),
		},
	}

	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu adalah assistant perekomendasi traveling itinerary destination",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu akan memilih destinasti mana saja yang sesuai dengan user inginkan dari pagi hingga malam",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `format waktunya akan seperti ini ["morning", "afternoon", "evening", "night"]`,
		},
		destinationData[0],
		destinationData[1],
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "destinasti yang saya berikan ini hanya destinasti dari negara singapura",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "pilihlah 2 hingga 4 destinasti id untuk setiap waktu-nya",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Saya ingin destinasi %s dengan trip %s", destination, trip),
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "hanya berikan waktu dan destination id nya saja",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: schema.ParseSchema(schemaGenerateItinerary),
		},
	}
}

var schemaGenerateItinerary = schema.Schema{
	Type: "array",
	Items: &schema.Schema{
		Type: "object",
		Properties: map[string]schema.Schema{
			"time": {
				Type:        "string",
				Description: `hanya berisi waktu ["morning", "afternoon", "evening", "night"]`,
			},
			"destination_ids": {
				Type:        "array",
				Description: "Jangan sampai ada destinasti id yang sama",
				Items: &schema.Schema{
					Type:        "number",
					Description: "hanya berisi destinasti id yang telah diberikan",
				},
			},
		},
	},
}
