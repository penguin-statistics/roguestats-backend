package repo

import (
	"context"

	"github.com/uptrace/bun"
	"go.uber.org/fx"

	"github.com/penguin-statistics/roguestats-backend/internal/model"
)

type User struct {
	fx.In

	DB *bun.DB
}

func (r *User) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	err := r.DB.NewSelect().
		Model(&user).
		Where("user_id = ?", id).
		Scan(ctx)

	return &user, err
}

func (r *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.DB.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)

	return &user, err
}

func (r *User) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.NewSelect().
		Model(&users).
		Scan(ctx)

	return users, err
}
