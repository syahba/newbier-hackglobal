package schema

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func GenerateItineraryMessage(destinations string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu akan membantu saya merancang itinerary destinasti dari waktu pagi hingga malam berdasarkan parameter-parameter yang saya berikan",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `Untuk waktunya sendiri hanya akan ada ["morning", "afternoon", "evening", "night"]`,
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Untuk parameternya akan saya berikan type aktifitas, type trip, hingga tanggal dan hari",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("untuk list destinasti nya sebagai berikut %s", destinations),
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "destinasti yang saya berikan ini hanya destinasti dari negara singapura",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "ingat kalau target kamu adalah orang-orang yang tinggal di singapura bisa jadi tourist atau warga lokal",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Pastikan setiap traveling pasti akan ada lunch dan dinner sehingga perlu tempat makan di waktu afternoon dan evening atau night",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "saya ingin itinerary dengan aktifitas Culinary Experiences, dengan trip Budget, pada hari minggu",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: ParseSchema(schemaGenerateItinerary),
		},
	}
}

var schemaGenerateItinerary = Schema{
	Type:        "array",
	Description: "Berikan itinerary destinasti yang diminta",
	Items: &Schema{
		Type: "object",
		Properties: map[string]Schema{
			"time": {
				Type:        "string",
				Description: "berisi waktu yang ditentukan",
			},
			"destinations": {
				Type:        "array",
				Description: "berikan id destinasti sesuai dengan yang telah saya berikan tadi",
				Items: &Schema{
					Type: "number",
				},
			},
		},
	},
}
