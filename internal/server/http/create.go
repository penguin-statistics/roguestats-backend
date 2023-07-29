package http

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/penguin-statistics/roguestats-backend/internal/app/appbundle"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "Penguin Stats: RogueStats",
		ServerHeader:          fmt.Sprintf("RogueStats/%s", appbundle.Version),
		StrictRouting:         true,
		CaseSensitive:         true,
		DisableStartupMessage: true,
		Immutable:             true,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: time.Second * 60,
	}))

	return app
}
