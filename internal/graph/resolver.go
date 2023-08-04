package graph

import (
	"github.com/uptrace/bun"
	"go.uber.org/fx"

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

	DB              *bun.DB
	AuthService     service.Auth
	UserService     service.User
	EventService    service.Event
	ResearchService service.Research
}
