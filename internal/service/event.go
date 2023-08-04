package service

import (
	"context"

	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/repo"
)

type Event struct {
	fx.In

	EventRepo repo.Event
}

func (s Event) CreateEvent(ctx context.Context, event *model.Event) error {
	return s.EventRepo.CreateEvent(ctx, event)
}

func (s Event) GetEvents(ctx context.Context) ([]*model.Event, error) {
	return s.EventRepo.GetEvents(ctx)
}
