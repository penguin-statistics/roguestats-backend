package middleware

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("middleware", fx.Populate(&Middleware{}))
}
