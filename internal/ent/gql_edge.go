// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
)

func (e *Event) User(ctx context.Context) (*User, error) {
	result, err := e.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryUser().Only(ctx)
	}
	return result, err
}

func (e *Event) Research(ctx context.Context) (*Research, error) {
	result, err := e.Edges.ResearchOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryResearch().Only(ctx)
	}
	return result, err
}
