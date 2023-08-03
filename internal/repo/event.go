package repo

import (
	"context"

	"github.com/uptrace/bun"
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/model"
)

type Event struct {
	fx.In

	DB *bun.DB
}

func (r *Event) CreateEvent(ctx context.Context, event *model.Event) error {
	_, err := r.DB.NewInsert().Model(event).Exec(ctx)
	return err
}

func (r *Event) GetEvents(ctx context.Context) ([]*model.Event, error) {
	var events []*model.Event

	err := r.DB.NewSelect().
		Model(&events).
		Scan(ctx)

	return events, err
}
