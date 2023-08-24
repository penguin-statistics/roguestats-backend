package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Meta struct {
	fx.In

	Route            fiber.Router `name:"internal"`
}

func RegisterMeta(c GraphQL) {
	c.Route.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
