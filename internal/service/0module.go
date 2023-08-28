package service

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("service",
		fx.Provide(NewJWT),
		fx.Provide(NewTurnstile))
}
