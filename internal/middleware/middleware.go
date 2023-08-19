package middleware

import (
	"strings"

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
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Next()
		}

		user, err := m.Auth.AuthByToken(ctx.Context(), strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
		}

		ctx.Context().SetUserValue("currentUser", user)

		return ctx.Next()
	}
}
