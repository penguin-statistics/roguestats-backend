package appcontext

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/penguin-statistics/roguestats-backend/internal/model"
)

var (
	ctxKeyCurrentUser = "currentUser"
	ctxKeyFiberCtx    = "fiberCtx"
)

// type ctxKey string

func CurrentUser(ctx context.Context) *model.User {
	v := ctx.Value(ctxKeyCurrentUser)
	if v == nil {
		return nil
	}
	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}

func WithCurrentUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, ctxKeyCurrentUser, user)
}

func FiberCtx(ctx context.Context) *fiber.Ctx {
	v := ctx.Value(ctxKeyFiberCtx)
	if v == nil {
		return nil
	}
	fiberCtx, ok := v.(*fiber.Ctx)
	if !ok {
		return nil
	}
	return fiberCtx
}

func WithFiberCtx(ctx context.Context, fiberCtx *fiber.Ctx) context.Context {
	return context.WithValue(ctx, ctxKeyFiberCtx, fiberCtx)
}
