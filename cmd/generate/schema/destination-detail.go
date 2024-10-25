package schema

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func DetailDestination(destinationName string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Kamu akan membantu saya untuk mencaritahu informasi tentang destinasti yang saya kasih ini",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Berikan saya deskripsi tentang destinasti tersebut yang relevan",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Berikan juga aktifitas ataupun produk bisa makanan atau barang apa saja yang paling rekomendasi di destinasti tersebut",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "destinasti yang saya berikan ini hanya destinasti dari negara singapura",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("bisa cari tahu informasi tentang destinasti %s?", destinationName),
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Harus berikan jawaban dengan bahasa inggris",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: ParseSchema(schemaDetailDestination),
		},
	}
}

type DestinationDetail struct {
	Description string   `json:"description"`
	Product     []string `json:"product"`
}

var schemaDetailDestination = Schema{
	Type:        "object",
	Description: "berikan informasi destinasti yang ada",
	Properties: map[string]Schema{
		"description": {
			Type:        "string",
			Description: "isi deskripsi singkat saja tentang destinasti tersebut",
		},
		"product": {
			Type:        "array",
			Description: "berisi aktifitas ataupun produk bisa makanan atau barang",
			Items: &Schema{
				Type:        "string",
				Description: "berikan jawaban dengna singkat padat dan jelas",
			},
		},
	},
}
