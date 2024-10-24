package schema

import (
	"encoding/json"
	"fmt"
)

type Type string

var (
	TypeString Type = "string"
	TypeObject Type = "object"
	TypeArray  Type = "array"
	TypeNumber Type = "number"
)

type Schema struct {
	Type        Type              `json:"type"`
	Description string            `json:"description"`
	Required    []string          `json:"required,omitempty"`
	Properties  map[string]Schema `json:"properties,omitempty"`
	Items       *Schema           `json:"items,omitempty"`
}

func ParseSchema(payload Schema) string {
	bp, _ := json.Marshal(payload)
	return fmt.Sprintf("Follow minified JSON schema.<JSONSchema>%s</JSONSchema>", string(bp))
}
