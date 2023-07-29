package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/app"
	"github.com/penguin-statistics/roguestats-backend/internal/app/appcontext"
)

func Start(module fx.Option) {
	app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
