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
	FindUsersByIds(ctx context.Context, ids ...string) ([]*model.User, error)
}

type UserRemover interface {
	RemoveUsers(ctx context.Context, users ...*model.User) error
}

type UserSaver interface {
	CreateUser(ctx context.Context, id uuid.UUID, password string, contacts ...*model.UserContact) (*model.User, error)
	UpdateUser(ctx context.Context, login string) (*model.User, error)
}
