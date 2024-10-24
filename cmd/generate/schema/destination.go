package schema

import (
	"github.com/sashabaranov/go-openai"
)

var AdditionalDestination = []openai.ChatCompletionMessage{
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Kamu akan membantu saya untuk menentukan destinasti yang saya kasih ini termasuk type aktifitas apa saja",
	},
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Kamu juga akan membantu saya untuk menentukan destinasti yang saya kasih ini termasuk type trip apa saja",
	},
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: "Kamu juga akan membantu saya untuk menentukan destinasti yang saya kasih ini memiliki produk atau jasa apa saja",
	},
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: `Berikut adalah tipe aktifitas nya ["sightseeing","culinary_experiences", "outdoor_activities", "adventure", "cultural_experiences", "shopping"]`,
	},
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: `Berikut adalah tipe trip nya ["relaxation", "budget", "luxury", "family_friendly", "adventure",]`,
	},
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: "destinasti yang saya berikan ini hanya destinasti dari negara singapura",
	},
	{
		Role:    openai.ChatMessageRoleUser,
		Content: "bisa cari tahu destinasti MINT Museum of Toys termasuk tipe aktifitas apa?",
	},
	{
		Role:    openai.ChatMessageRoleSystem,
		Content: ParseSchema(schemaAdditionalDestination),
	},
}

var schemaAdditionalDestination = Schema{
	Type:        "object",
	Description: "berikan data type aktifitas dan type trip berdasarkan destinasti tersebut",
	Properties: map[string]Schema{
		"activity": {
			Type:        "array",
			Description: "isi type aktifitas disini",
			Items: &Schema{
				Type: "string",
			},
		},
		"trip": {
			Type:        "array",
			Description: "isi type trip disini",
			Items: &Schema{
				Type: "string",
			},
		},
		"product": {
			Type:        "array",
			Description: "isi product disini",
			Items: &Schema{
				Type: "object",
				Properties: map[string]Schema{
					"name": {
						Type:        "string",
						Description: "berikan nama produk atau jasa",
					},
					"unit": {
						Type:        "string",
						Description: "ini bisa berupa unit atau jam dan lain-lain, tanpa imbuhan apapun",
					},
					"price": {
						Type:        "number",
						Description: "berikan harga produk atau jasa berdasarkan per unit nya",
					},
				},
			},
		},
	},
}
