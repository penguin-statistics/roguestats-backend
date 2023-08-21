package repo

import (
	"context"

	"github.com/uptrace/bun"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/model"
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

func (r *Event) GetEventsByResearchID(ctx context.Context, researchID string) ([]*model.Event, error) {
	var events []*model.Event

	err := r.DB.NewSelect().
		Model(&events).
		Where("research_id = ?", researchID).
		Scan(ctx)

	return events, err
}

func (r *Event) GetPaginatedEventsByResearchID(ctx context.Context, researchID string, first int, after string) ([]*model.Event, error) {
	var events []*model.Event

	query := r.DB.NewSelect().Model(&events)
	if researchID != "" {
		query = query.Where("research_id = ?", researchID)
	}
	if after != "" {
		query = query.Where("event_id > ?", after)
	}

	err := query.
		Order("event_id ASC").
		Limit(first).
		Scan(ctx)
	return events, err
}
