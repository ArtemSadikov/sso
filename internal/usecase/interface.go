package usecase

import (
	"context"

	"cinematic.com/sso/internal/domain/model"
)

type UseCases interface {
	AuthUseCase
}

type AuthUseCase interface {
	AuthByCredentials(ctx context.Context, login string, password string) (*AuthByCredentialsResultDto, error)
	RegisterUser(ctx context.Context, login string, password string) (*RegisterResultDto, error)
	ValidateToken(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, token string) (*model.Token, error)
}
