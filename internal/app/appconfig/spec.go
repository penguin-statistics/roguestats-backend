package appconfig

import (
	"time"

	"exusiai.dev/roguestats-backend/internal/app/appenv"
)

// ConfigSpec is the configuration specification.
type ConfigSpec struct {
	// DatabaseURL is the URL to the PostgreSQL database.
	DatabaseURL string `split_words:"true" required:"true"`

	// RedisURL is the URL to the Redis database.
	RedisURL string `split_words:"true" required:"true"`

	// ServiceListenAddress is the address that the Fiber HTTP server will listen on.
	ServiceListenAddress string `split_words:"true" required:"true" default:":3000"`

	// LogJSONStdout is the flag to enable JSON logging to stdout and disable logging to file.
	LogJSONStdout bool `split_words:"true" required:"true" default:"false"`

	// LogLevel is the log level. Valid values are: trace, debug, info, warn, error, fatal, panic.
	LogLevel ConfigLogLevel `split_words:"true" required:"true" default:"info"`

	// JWTPublicKey is the public key used to verify the JWT token.
	JWTPublicKey []byte `split_words:"true" required:"true"`

	// JWTPrivateKey is the private key used to sign the JWT token.
	JWTPrivateKey []byte `split_words:"true" required:"true"`

	// JWTExpiration is the expiration time of the JWT token. (default: 2 weeks (14 days, 336 hours))
	JWTExpiration time.Duration `split_words:"true" required:"true" default:"336h"`

	// JWTAutoRenewalTime is the time before expiration that a new JWT token will be
	// issued automatically to the client via the header that is the same as the
	// one used to authenticate. (default: 1 week (7 days, 168 hours))
	JWTAutoRenewalTime time.Duration `split_words:"true" required:"true" default:"168h"`

	// TurnstileSecret is the secret used to verify the turnstile response.
	TurnstileSecret string `split_words:"true" required:"true"`

	// ResendApiKey is the API key used to send emails via Resend.
	ResendApiKey string `split_words:"true" required:"true"`

	// PasswordResetTokenTTL is the time to live of the password reset token.
	PasswordResetTokenTTL time.Duration `split_words:"true" required:"true" default:"1h"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppEnv is the application context
	AppEnv appenv.Ctx
}
