package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/santhosh-tekuri/jsonschema/v5"

	"exusiai.dev/roguestats-backend/internal/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.User, error) {
	return r.AuthService.AuthByLoginInput(ctx, input)
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	// FIXME: should use a global snowflake node or something like an ID generator
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	user, err := r.AuthService.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	// get schema from research
	research, err := r.ResearchService.GetResearchByID(ctx, input.ResearchID)
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

	event := &model.Event{
		ID:         node.Generate().String(),
		ResearchID: input.ResearchID,
		Content:    input.Content,
		UserID:     user.ID,
		CreatedAt:  time.Now(),
		UserAgent:  input.UserAgent,
	}
	err = r.EventService.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// Me is the resolver for the `me` field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return r.AuthService.CurrentUser(ctx)
}

// Research is the resolver for the research field.
func (r *queryResolver) Research(ctx context.Context, id string) (*model.Research, error) {
	return r.ResearchService.GetResearchByID(ctx, id)
}

// Researches is the resolver for the researches field.
func (r *queryResolver) Researches(ctx context.Context) ([]*model.Research, error) {
	return r.ResearchService.GetAllResearch(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
