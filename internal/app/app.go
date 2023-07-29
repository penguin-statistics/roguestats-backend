package app

import (
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/app/appconfig"
	"github.com/penguin-statistics/roguestats-backend/internal/app/appenv"
	"github.com/penguin-statistics/roguestats-backend/internal/controller"
	"github.com/penguin-statistics/roguestats-backend/internal/infra"
	"github.com/penguin-statistics/roguestats-backend/internal/middleware"
	"github.com/penguin-statistics/roguestats-backend/internal/repo"
	"github.com/penguin-statistics/roguestats-backend/internal/server"
	"github.com/penguin-statistics/roguestats-backend/internal/service"
	"github.com/penguin-statistics/roguestats-backend/internal/x/logger"
	"github.com/penguin-statistics/roguestats-backend/internal/x/logger/fxlogger"
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
		repo.Module(),
		service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
