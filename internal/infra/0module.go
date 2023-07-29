package infra

import (
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/infra/db"
)

func Module() fx.Option {
	return fx.Module("infra", db.Module())
}
