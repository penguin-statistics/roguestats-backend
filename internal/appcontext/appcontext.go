package appcontext

import (
	"context"

	"exusiai.dev/roguestats-backend/internal/ent"
	"github.com/gofiber/fiber/v2"
)

var (
	CtxKeyCurrentUser ctxKey = "currentUser"
	CtxKeyFiberCtx    ctxKey = "fiberCtx"
)

type ctxKey string

func CurrentUser(ctx context.Context) *ent.User {
	v := ctx.Value(CtxKeyCurrentUser)
	if v == nil {
		return nil
	}
	u, ok := v.(*ent.User)
	if !ok {
		return nil
	}
	return u
}

func WithCurrentUser(ctx context.Context, user *ent.User) context.Context {
	return context.WithValue(ctx, CtxKeyCurrentUser, user)
}

func FiberCtx(ctx context.Context) *fiber.Ctx {
	v := ctx.Value(CtxKeyFiberCtx)
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
	return context.WithValue(ctx, CtxKeyFiberCtx, fiberCtx)
}
