package cli

import (
	"context"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/app"
	"exusiai.dev/roguestats-backend/internal/app/appenv"
)

func Start(module fx.Option) error {
	return app.New(appenv.Declare(appenv.EnvCLI), module).
		Start(context.Background())
}

func RunFunc(ctx *cli.Context, r any) error {
	return app.New(appenv.Declare(appenv.EnvCLI), fx.Supply(ctx), fx.Invoke(r)).
		Start(ctx.Context)
}
