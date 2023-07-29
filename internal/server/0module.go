package server

import (
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/server/http"
)

func Module() fx.Option {
	return fx.Module("server", http.Module())
}
