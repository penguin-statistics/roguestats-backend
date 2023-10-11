package service

import (
	"context"

	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/ent"
)

type QueryPreset struct {
	fx.In

	Ent *ent.Client
}

func (s QueryPreset) GetQueryPreset(ctx context.Context, queryPresetID string) (*ent.QueryPreset, error) {
	return s.Ent.QueryPreset.Get(ctx, queryPresetID)
}
