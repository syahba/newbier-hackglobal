package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Map struct{
	apiKey string
}

func getMap(Key string) *Map{
	return &Map{apiKey:Key}
}

func (gmap *Map) GetDistance(origin,destination string)(string,error){
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?destinations=%s&origins=%s&units=metric&key=%s",
		destination,
		origin,
		gmap.apiKey,
	)

	resp, err := http.Get(url)
    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "",fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "",fmt.Errorf("failed to read response body: %w", err)
    }

    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        return "", fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    rows, ok := result["rows"].([]interface{})
    if !ok || len(rows) == 0 {
        return "",errors.New("rows not found or empty")
    }

    elements, ok := rows[0].(map[string]interface{})["elements"].([]interface{})
    if !ok || len(elements) == 0 {
        return "",errors.New("elements not found or empty")
    }

    distance, ok := elements[0].(map[string]interface{})["distance"].(map[string]interface{})
    if !ok {
        return "",errors.New("distance not found")
    }

    distanceText, _ := distance["text"].(string)

    return distanceText, nil
}