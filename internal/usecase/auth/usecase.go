package auth

import (
	"context"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/domain/service"
	"cinematic.com/sso/internal/usecase"
	"cinematic.com/sso/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	logger   *slog.Logger
	userSrv  service.UserService
	tokenSrv service.TokenService
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
func (a *authUseCase) RegisterUser(ctx context.Context, login string, password string) (*usecase.RegisterResultDto, error) {
	op := "Auth.RegisterUser"

	log := a.logger.With(
		slog.String("op", op),
		slog.String("login", login),
	)

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.ErrorContext(ctx, "Error from generating pass", utils.LogErr(err))
		return nil, err
	}

	user, err := a.userSrv.CreateUser(ctx, string(pass), login)
	if err != nil {
		log.ErrorContext(ctx, "Error from saving user", utils.LogErr(err))
		return nil, err
	}

	tokens, err := a.tokenSrv.GeneratePair(ctx, user)
	if err != nil {
		log.ErrorContext(ctx, "Error from generating tokens", utils.LogErr(err))
		return nil, err
	}

	return &usecase.RegisterResultDto{
		User:         user,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (a *authUseCase) ValidateToken(ctx context.Context, token string) error {
	op := "Auth.ValidateToken"

	log := a.logger.With(slog.String("op", op))
	t := model.NewToken(token, nil)

	if err := a.tokenSrv.Validate(ctx, t); err != nil {
		log.ErrorContext(ctx, "Error from validating token", utils.LogErr(err))
		return err
	}

	return nil
}

func (a *authUseCase) RefreshToken(ctx context.Context, token string, userId uuid.UUID) (*model.Token, error) {
	op := "Auth.ValidateToken"

	log := a.logger.With(
		slog.String("op", op),
		slog.String("uid", userId.String()),
	)

  user, err := a.userSrv.FindUsersByIds(ctx, &userId)
  if err != nil {
		log.ErrorContext(ctx, "Error from searching user", utils.LogErr(err))
    return nil, err
  }

  res, err := a.tokenSrv.RefreshToken(ctx, user[0], model.NewToken(token, nil))
  if err != nil {
    log.ErrorContext(ctx, "Error from refreshing token", utils.LogErr(err))
    return nil, err
  }

  return res, nil
}

func NewAuthUseCase(
	logger *slog.Logger,
	userSrv service.UserService,
	tokenSrv service.TokenService,
) *authUseCase {
	return &authUseCase{logger, userSrv, tokenSrv}
}
