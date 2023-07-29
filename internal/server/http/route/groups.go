package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type RouteGroups struct {
	fx.Out

	Root     fiber.Router `name:"root"`
	Internal fiber.Router `name:"internal"`
}

func CreateGroups(app *fiber.App) RouteGroups {
	return RouteGroups{
		Root:     app.Group("/"),
		Internal: app.Group("/_"),
	}
}
