package service

import (
	"context"

	"github.com/penguin-statistics/roguestats-backend/internal/model"
	"github.com/penguin-statistics/roguestats-backend/internal/repo"
)

type Post struct {
	postRepo *repo.Post
}

func NewPost(postRepo *repo.Post) *Post {
	return &Post{
		postRepo: postRepo,
	}
}

func (s *Post) GetPosts(ctx context.Context) ([]model.Post, error) {
	posts, err := s.postRepo.GetPosts(ctx)

	return posts, err
}
