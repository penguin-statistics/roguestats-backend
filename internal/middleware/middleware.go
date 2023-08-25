package middleware

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/service"
)

type Middleware struct {
	fx.In

	Auth service.Auth
}

func (m Middleware) CurrentUser() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get(fiber.HeaderAuthorization)
		if token == "" {
			// some routes don't require authentication so we just skip them
			return ctx.Next()
		}

		user, err := m.Auth.AuthByToken(ctx.Context(), strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			spew.Dump(err)
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
		}

		ctx.Context().SetUserValue("currentUser", user)

		return ctx.Next()
	}
}
