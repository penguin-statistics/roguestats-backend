package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"exusiai.dev/roguestats-backend/internal/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.User, error) {
	return r.AuthService.AuthByLoginInput(ctx, input)
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	return r.EventService.CreateEventFromInput(ctx, &input)
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

// GroupCount is the resolver for the groupCount field.
func (r *queryResolver) GroupCount(ctx context.Context, input model.GroupCountInput) (*model.GroupCountResult, error) {
	groupCountResult, err := r.EventService.CalculateStats(ctx, input.FilterInput, input.ResultMappingInput)
	if err != nil {
		return nil, err
	}
	return groupCountResult, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
