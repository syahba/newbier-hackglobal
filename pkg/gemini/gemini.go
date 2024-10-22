package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GenerativeAi struct {
	Model *genai.GenerativeModel
}

func GetModel(apiKey string) *GenerativeAi {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	genAi := &GenerativeAi{
		Model: client.GenerativeModel("gemini-1.5-flash"),
	}

	return genAi
}

func getText(resp *genai.GenerateContentResponse) string {
	result := ""
	
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%v",part)
			}
		}
	}

	return result
}

func (genAi *GenerativeAi)Chat(input string) (string,error){
	ctx := context.Background()
	resp,err := genAi.Model.GenerateContent(ctx,genai.Text(input))

	if err!=nil{
		return "",err
	}
	return getText(resp),nil
}