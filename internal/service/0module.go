package service

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("service",
		fx.Populate(&Auth{}),
		fx.Populate(&User{}),
		fx.Provide(NewJWT),
		fx.Provide(NewTurnstile),
		fx.Populate(&Event{}),
		fx.Populate(&Directive{}))
}
