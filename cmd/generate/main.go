package main

import (
	"encoding/json"
	"fmt"
	"log"
	"newbier-hackglobal/cmd/generate/schema"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"
	"newbier-hackglobal/pkg/database/model"
	"os"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load database: %v", err))
	}
	ai := chatgpt.GetModel(cfg.ChatGPTKey)

	// var (
	// 	activity = "sightseeing"
	// 	trip     = "budget"
	// )

	// table := model.Destination{}
	// destinations := []map[string]any{}
	// db.Select(table.ColumnName("id"), table.ColumnName("name"), table.ColumnName("type")).
	// 	Model(table).
	// 	Joins(`JOIN "destination_parameters" dp1 ON `+table.ColumnName("id")+` = dp1.destination_id AND dp1."name" = ?`, activity).
	// 	Joins(`JOIN "destination_parameters" dp2 ON `+table.ColumnName("id")+` = dp2.destination_id AND dp2."name" = ?`, trip).
	// 	Find(&destinations)

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// r.Shuffle(len(destinations), func(i, j int) { destinations[i], destinations[j] = destinations[j], destinations[i] })

	// result, _ := ai.GenerateCustom(0.5, 700, schema.GenerateItinerary(destinations, activity, trip))
	// fmt.Println(result)

	generateJsonSeeder[schema.DestinationDetail](ai, db, schema.DetailDestination, "destination-additional/destination_detail2")
	// generateJsonSeeder[schema.DestinationAdditional](ai, db, schema.AdditionalDestination, "destination-additional/destination_additional2")
}

type FuncMessage func(destinationName string) []openai.ChatCompletionMessage

func generateJsonSeeder[T any](ai *chatgpt.Model, db *gorm.DB, message FuncMessage, name string) {
	start := time.Now()

	file, err := os.Create(fmt.Sprintf("cmd/seeder/data/%s.json", name))
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	file.WriteString("[")

	defer func() {
		file.WriteString("]")
		file.Close()
		elapsed := time.Since(start)
		fmt.Printf("Program selesai! Waktu eksekusi: %v\n", elapsed)
	}()

	var destinations = make([]*model.Destination, 0)
	db.Offset(257).Order("id ASC").Find(&destinations)

	for idx, elm := range destinations {
		result, _ := ai.Generate(message(elm.Name))
		// if err != nil {
		// 	time.Sleep(10 * time.Minute)
		// }

		cleanedInput := strings.Trim(result, "```json")
		cleanedInput = strings.Trim(cleanedInput, "```")

		var destinationDetail = new(T)
		json.Unmarshal([]byte(cleanedInput), destinationDetail)

		jsonData, err := json.Marshal(schema.Response[*T]{
			ID:   int(elm.ID),
			Data: destinationDetail,
		})
		if err != nil {
			log.Fatalf("Error encoding JSON: %v", err)
		}
		file.Write(jsonData)

		if idx < len(destinations)-1 {
			file.WriteString(",")
		}

		log.Printf("Data: %s success", elm.Name)
		time.Sleep(25 * time.Second)
	}

}
