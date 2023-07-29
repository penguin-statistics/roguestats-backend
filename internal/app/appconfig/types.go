package appconfig

import (
	"encoding/base64"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type ConfigLogLevel zerolog.Level

// ensure ConfigLogLevel implements envconfig.Decoder
var _ envconfig.Decoder = (*ConfigLogLevel)(nil)

func (c *ConfigLogLevel) Decode(value string) error {
	level, err := zerolog.ParseLevel(value)
	if err != nil {
		return err
	}

	*c = ConfigLogLevel(level)
	return nil
}

type Base64EncodedJWTSecret []byte

// ensure Base64EncodedJWTSecret implements envconfig.Decoder
var _ envconfig.Decoder = (*Base64EncodedJWTSecret)(nil)

func (c *Base64EncodedJWTSecret) Decode(value string) (err error) {
	*c, err = base64.StdEncoding.DecodeString(value)
	return err
}
