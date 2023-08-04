package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"exusiai.dev/roguestats-backend/internal/app/appbundle"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:                 "Penguin Stats: RogueStats",
		ServerHeader:            fmt.Sprintf("RogueStats/%s", appbundle.Version),
		StrictRouting:           true,
		CaseSensitive:           true,
		DisableStartupMessage:   true,
		Immutable:               true,
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"::1", "127.0.0.1", "10.0.0.0/8"},
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	return app
}
