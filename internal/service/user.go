package service

import (
	"context"

	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/model"
	"github.com/penguin-statistics/roguestats-backend/internal/repo"
)

type User struct {
	fx.In

	UserRepo repo.User
}

func (s User) GetUsers(ctx context.Context) ([]*model.User, error) {
	return s.UserRepo.GetUsers(ctx)
}
