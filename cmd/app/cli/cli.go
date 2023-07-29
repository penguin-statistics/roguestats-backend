package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/app"
	"github.com/penguin-statistics/roguestats-backend/internal/app/appenv"
)

func Start(module fx.Option) {
	app.New(appenv.Declare(appenv.EnvCLI), module).Start(context.Background())
}
