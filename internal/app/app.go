package app

import (
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
	"exusiai.dev/roguestats-backend/internal/app/appenv"
	"exusiai.dev/roguestats-backend/internal/controller"
	"exusiai.dev/roguestats-backend/internal/infra"
	"exusiai.dev/roguestats-backend/internal/middleware"
	"exusiai.dev/roguestats-backend/internal/server"
	"exusiai.dev/roguestats-backend/internal/service"
	"exusiai.dev/roguestats-backend/internal/x/logger"
	"exusiai.dev/roguestats-backend/internal/x/logger/fxlogger"
)

func New(ctx appenv.Ctx, additionalOpts ...fx.Option) *fx.App {
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
		fx.Supply(conf),
		middleware.Module(),
		controller.Module(),
		infra.Module(),
		service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
