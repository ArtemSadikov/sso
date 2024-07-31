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
func (u *userService) CreateUser(ctx context.Context, password, login string) (*model.User, error) {
	return u.repo.CreateUser(ctx, login, password)
}

// DeleteUsers implements UserService.
func (u *userService) DeleteUsers(ctx context.Context, users ...*model.User) error {
	return nil
}

// FindUsersByIds implements UserService.
func (u *userService) FinUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return u.repo.FindUserById(ctx, id)
}

func (u *userService)	FindUserByLogin(ctx context.Context, login string) (*model.User, error) {
	return u.repo.FindUserByLogin(ctx, login)
}

// UpdateUsers implements UserService.
func (u *userService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func NewUserService(logger *slog.Logger, repo repository.UserRepository) *userService {
	return &userService{logger, repo}
}
