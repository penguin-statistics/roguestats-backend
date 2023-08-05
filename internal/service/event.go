package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/repo"
)

type Event struct {
	fx.In

	EventRepo       repo.Event
	ResearchService Research
	AuthService     Auth
}

func (s Event) CreateEvent(ctx context.Context, event *model.Event) error {
	return s.EventRepo.CreateEvent(ctx, event)
}

func (s Event) GetEvents(ctx context.Context) ([]*model.Event, error) {
	return s.EventRepo.GetEvents(ctx)
}

func (s Event) CreateEventFromInput(ctx context.Context, input *model.NewEvent) (*model.Event, error) {
	event, err := s.convertFromEventInputToEvent(ctx, input)
	if err != nil {
		return nil, err
	}
	err = s.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s Event) convertFromEventInputToEvent(ctx context.Context, input *model.NewEvent) (*model.Event, error) {
	user, err := s.AuthService.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	// get schema from research
	research, err := s.ResearchService.GetResearchByID(ctx, input.ResearchID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("research not found")
		}
		return nil, err
	}

	// validate event json
	schema, err := json.Marshal(research.Schema)
	if err != nil {
		return nil, err
	}
	sch, err := jsonschema.CompileString("schema.json", string(schema))
	if err != nil {
		return nil, err
	}
	if err = sch.Validate(input.Content); err != nil {
		return nil, err
	}

	// FIXME: should use a global snowflake node or something like an ID generator
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	event := &model.Event{
		ID:         node.Generate().String(),
		ResearchID: input.ResearchID,
		Content:    input.Content,
		UserID:     user.ID,
		CreatedAt:  time.Now(),
		UserAgent:  input.UserAgent,
	}
	return event, nil
}
