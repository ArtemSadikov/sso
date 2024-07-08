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
	CreateUsers(ctx context.Context, users ...*model.User) ([]*model.User, error)
	UpdateUsers(ctx context.Context, users ...*model.User) (*model.User, error)
	DeleteUsers(ctx context.Context, users ...*model.User) error
	FindUsersByIds(ctx context.Context, ids ...*uuid.UUID) ([]*model.User, error)
}
