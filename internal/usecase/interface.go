package usecase

import (
	"context"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type UseCases interface {
	AuthUseCase
}

type AuthUseCase interface {
	AuthByCredentials(ctx context.Context, login string, password string) (*model.User, error)
	RegisterUser(ctx context.Context, login string, password string) (*RegisterResultDto, error)
	ValidateToken(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, token string, userId uuid.UUID) (*model.Token, error)
}
