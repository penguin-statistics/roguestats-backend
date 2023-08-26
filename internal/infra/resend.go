package infra

import (
	"github.com/resendlabs/resend-go"

	"exusiai.dev/roguestats-backend/internal/app/appconfig"
)

func Resend(conf *appconfig.Config) *resend.Client {
	return resend.NewClient(conf.ResendApiKey)
}
