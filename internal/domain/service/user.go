package service

import (
	"context"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type userService struct {
	logger *slog.Logger
}

// CreateUsers implements UserService.
func (u *userService) CreateUsers(ctx context.Context, users ...*model.User) ([]*model.User, error) {
	return nil, nil
}

// DeleteUsers implements UserService.
func (u *userService) DeleteUsers(ctx context.Context, users ...*model.User) error {
	return nil
}

// FindUsersByIds implements UserService.
func (u *userService) FindUsersByIds(ctx context.Context, ids ...*uuid.UUID) ([]*model.User, error) {
	return nil, nil
}

// UpdateUsers implements UserService.
func (u *userService) UpdateUsers(ctx context.Context, users ...*model.User) (*model.User, error) {
	return nil, nil
}

func NewUserService(logger *slog.Logger) *userService {
	return &userService{logger}
}
