package service

import (
	"context"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type Service interface {
	UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUsers(ctx context.Context, users ...*model.User) error
	FindUsersByIds(ctx context.Context, ids ...*uuid.UUID) ([]*model.User, error)
}
