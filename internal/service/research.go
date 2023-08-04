package service

import (
	"context"

	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/repo"
)

type Research struct {
	fx.In

	ResearchRepo repo.Research
}

func (s Research) GetAllResearch(ctx context.Context) ([]*model.Research, error) {
	return s.ResearchRepo.GetAllResearch(ctx)
}

func (s Research) GetResearchByID(ctx context.Context, id string) (*model.Research, error) {
	return s.ResearchRepo.GetResearchByID(ctx, id)
}
