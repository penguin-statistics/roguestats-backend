package model

type Research struct {
	ID     string                 `json:"id" bun:"research_id"`
	Name   string                 `json:"name"`
	Schema map[string]interface{} `json:"schema"`
}
