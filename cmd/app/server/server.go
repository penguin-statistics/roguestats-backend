package server

import (
	"context"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/app"
	"exusiai.dev/roguestats-backend/internal/app/appconfig"
	"exusiai.dev/roguestats-backend/internal/app/appenv"
)

func Run() {
	app.New(appenv.Declare(appenv.EnvServer), fx.Invoke(run)).Run()
}

func run(lc fx.Lifecycle, app *fiber.App, conf *appconfig.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", conf.ServiceListenAddress)
			if err != nil {
				return err
			}
			log.Info().Str("address", conf.ServiceListenAddress).Msg("server started")

			go func() {
				if err := app.Listener(ln); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("gracefully shutting down server")
			if err := app.Shutdown(); err != nil {
				log.Error().Err(err).Msg("error occurred while gracefully shutting down server")
				return err
			}
			log.Info().Msg("graceful server shut down completed")
			return nil
		},
	})
}
