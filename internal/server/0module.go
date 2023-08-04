package server

import (
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/server/http"
)

func Module() fx.Option {
	return fx.Module("server", http.Module())
}
