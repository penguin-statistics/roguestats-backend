// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CategoryCount struct {
	Category interface{} `json:"category"`
	Count    int         `json:"count"`
}

type Event struct {
	ID         string                 `json:"id" bun:"event_id"`
	ResearchID string                 `json:"research_id" bun:"research_id"`
	Content    map[string]interface{} `json:"content"`
	UserID     string                 `json:"user_id" bun:"user_id"`
	CreatedAt  time.Time              `json:"created_at"`
	UserAgent  *string                `json:"user_agent,omitempty"`
}

type GroupCountInput struct {
	ResearchID         string `json:"research_id"`
	FilterInput        string `json:"filterInput"`
	ResultMappingInput string `json:"resultMappingInput"`
}

type GroupCountResult struct {
	Results []*CategoryCount `json:"results"`
	Total   int              `json:"total"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewEvent struct {
	Content    map[string]interface{} `json:"content"`
	ResearchID string                 `json:"research_id"`
	UserAgent  *string                `json:"user_agent,omitempty"`
}

type Research struct {
	ID     string                 `json:"id" bun:"research_id"`
	Name   string                 `json:"name"`
	Schema map[string]interface{} `json:"schema"`
}

type Topic struct {
	ID                 *string `json:"id,omitempty" bun:"topic_id"`
	FilterInput        string  `json:"filterInput"`
	ResultMappingInput string  `json:"resultMappingInput"`
}

type User struct {
	ID         string                 `json:"id" bun:"user_id"`
	Name       string                 `json:"name"`
	Email      *string                `json:"email,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// User's encrypted credential
	Credential string `json:"-"`
}
