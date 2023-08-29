package userimport

import (
	appcli "exusiai.dev/roguestats-backend/cmd/app/cli"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "user-import",
		Usage: "Import users from a JSON file.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Usage:    "Token to use for authentication.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "File containing the users to import.",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return appcli.RunFunc(c, Run)
		},
	}
}
