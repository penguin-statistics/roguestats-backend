package script

import (
	"log"

	"exusiai.dev/roguestats-backend/cmd/app/cli/script/syncschema"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "script",
		Usage: "run scripts",
		Subcommands: []*cli.Command{
			{
				Name:  "import",
				Usage: "import data from temp csv",
				Action: func(c *cli.Context) error {
					importType := c.Args().Get(0)
					path := c.Args().Get(1)
					if path == "" {
						return cli.Exit("Please provide path to csv file", 1)
					}
					log.Default().Printf("importing %s data from '%s'\n", importType, path)
					if importType == "battle" {
						err := NewBattleCSVImport(path).Run()
						if err != nil {
							return err
						}
					} else if importType == "incident" {
						err := NewIncidentCSVImport(path).Run()
						if err != nil {
							return err
						}
					} else if importType == "rest" {
						err := NewRestCSVImport(path).Run()
						if err != nil {
							return err
						}
					} else if importType == "portal" {
						err := NewPortalCSVImport(path).Run()
						if err != nil {
							return err
						}
					} else {
						return cli.Exit("Unknown import type", 1)
					}
					log.Default().Printf("finished importing %s data from '%s'\n", importType, path)
					return nil
				},
			},
			syncschema.Command(),
		},
	}
}
