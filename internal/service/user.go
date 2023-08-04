package service

import (
	"context"

	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/repo"
)

type User struct {
	fx.In

	UserRepo repo.User
}

func (s User) GetUsers(ctx context.Context) ([]*model.User, error) {
	return s.UserRepo.GetUsers(ctx)
}
