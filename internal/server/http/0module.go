package http

import (
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/server/http/route"
)

func Module() fx.Option {
	return fx.Module("http", fx.Provide(Create), route.Module())
}
