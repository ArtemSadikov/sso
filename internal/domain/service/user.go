package service

import (
	"context"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/infrastructure/storage/repository"
	"github.com/google/uuid"
)

type userService struct {
	logger *slog.Logger
	repo repository.UserRepository
}

// CreateUsers implements UserService.
func (u *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return u.repo.CreateUser(ctx, user.Login, user.Password)
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
func (u *userService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func NewUserService(logger *slog.Logger, repo repository.UserRepository) *userService {
	return &userService{logger, repo}
}
