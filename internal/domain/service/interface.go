package service

import (
	"context"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type Service interface {
	UserService
	TokenService
}

type UserService interface {
	CreateUser(ctx context.Context, password string, login string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUsers(ctx context.Context, users ...*model.User) error
	FinUserById(ctx context.Context, id uuid.UUID) (*model.User, error)
	FindUserByLogin(ctx context.Context, login string) (*model.User, error)
}

type TokenService interface {
	GeneratePair(ctx context.Context, user *model.User) (*TokenPair, error)
	RefreshToken(ctx context.Context, user *model.User, token *model.Token) (*model.Token, error)
	Validate(ctx context.Context, token *model.Token) error
}
