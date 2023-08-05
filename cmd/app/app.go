package app

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"exusiai.dev/roguestats-backend/cmd/app/cli/db"
	"exusiai.dev/roguestats-backend/cmd/app/cli/script"
	"exusiai.dev/roguestats-backend/cmd/app/server"
)

func Run() {
	app := &cli.App{
		Name: "app",
		Commands: []*cli.Command{
			server.Command(),
			db.Command(),
			script.Command(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}
}
