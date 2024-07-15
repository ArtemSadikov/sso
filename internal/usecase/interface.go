package usecase

import (
	"context"

	"cinematic.com/sso/internal/domain/model"
)

type UseCases interface {
	AuthUseCase
}

type AuthUseCase interface {
	AuthByCredentials(ctx context.Context, login string, password string) (*model.User, error)
	RegisterUser(ctx context.Context, login string, password string) (*model.User, error)
}
