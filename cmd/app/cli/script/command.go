package script

import (
	"github.com/urfave/cli/v2"

	"exusiai.dev/roguestats-backend/cmd/app/cli/script/syncschema"
	"exusiai.dev/roguestats-backend/cmd/app/cli/script/tempcsvimport"
	"exusiai.dev/roguestats-backend/cmd/app/cli/script/userimport"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "script",
		Usage: "run scripts",
		Subcommands: []*cli.Command{
			tempcsvimport.Command(),
			syncschema.Command(),
			userimport.Command(),
		},
	}
}
