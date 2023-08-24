package repo

import (
	"context"

	"github.com/uptrace/bun"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/model"
)

type Research struct {
	fx.In

	DB *bun.DB
}

func (r *Research) GetAllResearch(ctx context.Context) ([]*model.Research, error) {
	var research []*model.Research

	err := r.DB.NewSelect().
		Model(&research).
		Order("research_id ASC").
		Scan(ctx)

	return research, err
}

func (r *Research) GetResearchByID(ctx context.Context, id string) (*model.Research, error) {
	var research model.Research

	err := r.DB.NewSelect().
		Model(&research).
		Where("research_id = ?", id).
		Scan(ctx)

	return &research, err
}
