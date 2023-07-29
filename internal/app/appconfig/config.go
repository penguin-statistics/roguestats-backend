package appconfig

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"

	"github.com/penguin-statistics/roguestats-backend/internal/app/appenv"
)

func Parse(ctx appenv.Ctx) (*Config, error) {
	var conf ConfigSpec
	if err := envconfig.Process("roguestats", &conf); err != nil {
		return nil, err
	}

	return &Config{
		ConfigSpec: conf,
		AppEnv:     ctx,
	}, nil
}
