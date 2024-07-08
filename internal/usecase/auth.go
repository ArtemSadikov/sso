package usecase

import (
	"context"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/domain/service"
)

type authUseCase struct {
	logger  *slog.Logger
	userSrv service.UserService
}

// AuthByCredentials implements AuthUseCase.
func (a *authUseCase) AuthByCredentials(ctx context.Context) (*model.User, error) {
	users, err := a.userSrv.CreateUsers(ctx, model.NewUser(""))
  if err != nil {
    a.logger.Warn("Failed to auth")
    return nil, err
  }

  user := users[0]
  return user, nil
}

// RegisterUser implements AuthUseCase.
func (a *authUseCase) RegisterUser(ctx context.Context) (*model.User, error) {
	users, err := a.userSrv.CreateUsers(ctx, model.NewUser(""))
  if err != nil {
    a.logger.Warn("Failed to register")
    return nil, err
  }

  user := users[0]
  return user, nil
}

func NewAuthUseCase(
	logger *slog.Logger,
	userSrv service.UserService,
) *authUseCase {
	return &authUseCase{logger, userSrv}
}
