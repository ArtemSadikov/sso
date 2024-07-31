package repository

import (
	"context"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type Repository interface{}

type UserRepository interface {
	UserSaver
	UserSearcher
	UserRemover
}

type UserSearcher interface {
	FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
	FindUserByLogin(ctx context.Context, login string) (*model.User, error)
}

type UserRemover interface {
	RemoveUsers(ctx context.Context, users ...*model.User) error
}

type UserSaver interface {
	CreateUser(ctx context.Context, login, password string, contacts ...*model.UserContact) (*model.User, error)
	UpdateUser(ctx context.Context, login string) (*model.User, error)
}
