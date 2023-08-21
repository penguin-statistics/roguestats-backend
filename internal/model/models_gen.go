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
	ResearchID string                 `json:"researchId" bun:"research_id"`
	Content    map[string]interface{} `json:"content"`
	UserID     string                 `json:"userId" bun:"user_id"`
	CreatedAt  time.Time              `json:"createdAt"`
	UserAgent  *string                `json:"userAgent,omitempty"`
}

type EventsConnection struct {
	Edges    []*EventsEdge `json:"edges"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type EventsEdge struct {
	Node   *Event `json:"node"`
	Cursor string `json:"cursor"`
}

type GroupCountInput struct {
	ResearchID         string `json:"researchId"`
	FilterInput        string `json:"filterInput"`
	ResultMappingInput string `json:"resultMappingInput"`
}

type GroupCountResult struct {
	Results []*CategoryCount `json:"results"`
	Total   int              `json:"total"`
}

type LoginInput struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	TurnstileResponse string `json:"turnstileResponse"`
}

type NewEvent struct {
	Content    map[string]interface{} `json:"content"`
	ResearchID string                 `json:"researchId"`
	UserAgent  *string                `json:"userAgent,omitempty"`
}

type PageInfo struct {
	HasNextPage *bool  `json:"hasNextPage,omitempty"`
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
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
