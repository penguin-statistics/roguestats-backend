package service

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/appcontext"
)

type Directive struct {
	fx.In

	User User
}

func (s Directive) Private(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	user := appcontext.CurrentUser(ctx)
	log.Debug().Interface("user", user).Msg("directive: private")
	return next(ctx)
}
