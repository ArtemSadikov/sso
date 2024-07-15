package usecase

import (
	"context"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/domain/service"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	logger  *slog.Logger
	userSrv service.UserService
}

// AuthByCredentials implements AuthUseCase.
func (a *authUseCase) AuthByCredentials(ctx context.Context, login string, password string) (*model.User, error) {
	// users, err := a.userSrv.CreateUsers(ctx, model.NewUser(""))
  // if err != nil {
  //   a.logger.Warn("Failed to auth")
  //   return nil, err
  // }

  // user := users[0]
  return nil, nil
}

// RegisterUser implements AuthUseCase.
func (a *authUseCase) RegisterUser(ctx context.Context, login string, password string) (*model.User, error) {
  op := "Auth.RegisterUser"

  log := a.logger.With(
    slog.String("op", op),
    slog.String("login", login),
  )

  pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    log.ErrorContext(ctx, "Error from generating pass", slog.Any("error", slog.AnyValue(err)))
    return nil, err
  }

  user := model.NewUserWithPassword(login, string(pass))

  res, err := a.userSrv.CreateUser(ctx, user)
  if err != nil {
    log.ErrorContext(ctx, "Error from saving user", slog.Any("error", slog.AnyValue(err)))
    return nil, err
  }

  return res, nil
}

func NewAuthUseCase(
	logger *slog.Logger,
	userSrv service.UserService,
) *authUseCase {
	return &authUseCase{logger, userSrv}
}
