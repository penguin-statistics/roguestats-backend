package syncschema

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/ent/research"
)

type SyncSchemaCommandDeps struct {
	fx.In

	Ent *ent.Client
}

func Run(c *cli.Context, d SyncSchemaCommandDeps) error {
	tx, err := d.Ent.BeginTx(c.Context, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dir := c.String("dir")
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}

		// id is the beginning of the filename until the first dot.
		segments := strings.Split(info.Name(), ".")
		id := segments[0]
		name := segments[1]
		log.Info().Str("path", path).Str("id", id).Msg("processing research")

		jsonBytes, err := minifiedJsonFile(path)
		if err != nil {
			return err
		}

		// Check if the research exists.
		research, err := tx.Research.Query().Where(research.ID(id)).Only(c.Context)
		if err != nil {
			if ent.IsNotFound(err) {
				log.Info().Str("id", id).Msg("research does not exist, creating")
				// Create the research.
				research, err = tx.Research.Create().
					SetID(id).
					SetName(name).
					SetSchema(jsonBytes).
					Save(c.Context)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		// Update the research schema.
		log.Info().Str("id", id).Msg("updating research schema")
		_, err = research.Update().
			SetName(name).
			SetSchema(jsonBytes).
			Save(c.Context)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	log.Info().Msg("committing transaction")

	return tx.Commit()
}

func minifiedJsonFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
