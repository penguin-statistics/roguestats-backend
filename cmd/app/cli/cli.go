package cli

import (
	"context"

	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/app"
	"exusiai.dev/roguestats-backend/internal/app/appenv"
)

func Start(module fx.Option) {
	app.New(appenv.Declare(appenv.EnvCLI), module).Start(context.Background())
}
