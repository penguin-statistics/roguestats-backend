package infra

import (
	"github.com/resendlabs/resend-go"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
)

func Resend(conf *appconfig.Config) *resend.Client {
	client := resend.NewClient(conf.ResendApiKey)
	return client
}
