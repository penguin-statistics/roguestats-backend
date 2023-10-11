package syncschema

import (
	"github.com/urfave/cli/v2"

	appcli "exusiai.dev/roguestats-backend/cmd/app/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "sync-schema",
		Usage: "Sync JSON Schemas under DIR with the database.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "Directory containing the JSON Schemas.",
				Value:   "./schema/researches",
			},
		},
		Action: func(c *cli.Context) error {
			return appcli.RunFunc(c, Run)
		},
	}
}
