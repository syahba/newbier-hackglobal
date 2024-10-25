package schema

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func AdditionalDestination(destinationName string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
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
			Content: "Kamu juga akan membantu saya untuk menentukan destinasti yang saya kasih ini memiliki produk berupa barang atau jasa (seperti peminjaman barang atau yang lainnya) apa saja",
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `Berikut adalah tipe aktifitas nya ["sightseeing","culinary_experiences", "outdoor_activities", "adventure", "cultural_experiences", "shopping"]`,
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `Berikut adalah tipe trip nya ["relaxation", "budget", "luxury", "family_friendly", "adventure"]`,
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `Untuk produk dia akan memiliki data nama, harga dalam dolar singapura dan unit yang akan berisi item jika dia barang dan hour jika dia jasa`,
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: `Tidak semua harus destinasti ada produk nya`,
		},
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "destinasti yang saya berikan ini hanya destinasti dari negara singapura",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("bisa cari tahu tentang destinasti %s termasuk tipe aktifitas apa?", destinationName),
		},
		// {
		// 	Role:    openai.ChatMessageRoleSystem,
		// 	Content: "Harus berikan jawaban dengan bahasa inggris",
		// },
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: ParseSchema(schemaAdditionalDestination),
		},
	}
}

type DestinationAdditional struct {
	Activity []string  `json:"activity"`
	Trip     []string  `json:"trip"`
	Product  []Product `json:"product"`
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
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
