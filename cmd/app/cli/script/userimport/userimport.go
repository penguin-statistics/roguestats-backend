package userimport

import (
	"context"
	"encoding/csv"
	"os"

	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/service"
	"github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

type UserImportCommandDeps struct {
	fx.In

	Ent  *ent.Client
	Auth service.Auth
}

func Run(c *cli.Context, d UserImportCommandDeps) error {
	file := c.String("file")
	log.Info().Msgf("opening file %s", file)

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}
	rows = rows[1:]

	log.Info().Msg("starting to import users")

	for _, row := range rows {
		log.Info().Msgf("importing user %s", row[3])

		input := model.CreateUserInput{
			Name:  row[3],
			Email: row[2],
			Attributes: map[string]any{
				"contact_qq": row[4],
			},
		}
		err = CreateUser(c.String("token"), input)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("finished importing users")

	return nil
}

func CreateUser(token string, input model.CreateUserInput) error {
	client := graphql.NewClient("https://rogue.penguin-stats.io/graphql")
	req := graphql.NewRequest(`
	mutation CreateUserMutation($input: CreateUserInput!) {
		createUser(input: $input) {
			id
		}
	}`,
	)
	req.Var("input", input)
	req.Header.Set("Authorization", "Bearer "+token)
	var respData any

	return client.Run(context.Background(), req, &respData)
}
