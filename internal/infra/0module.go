package infra

import (
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module(), fx.Provide(Resend))
}
