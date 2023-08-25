package graph

import (
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/service"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ResolverDeps
}

type ResolverDeps struct {
	fx.In

	DB              *ent.Client
	AuthService     service.Auth
	UserService     service.User
	EventService    service.Event
	ResearchService service.Research
}
