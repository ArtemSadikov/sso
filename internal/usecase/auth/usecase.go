package auth

import (
	"context"
	"errors"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/domain/service"
	"cinematic.com/sso/internal/usecase"
	"cinematic.com/sso/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	logger   *slog.Logger
	userSrv  service.UserService
	tokenSrv service.TokenService
}

// AuthByCredentials implements AuthUseCase.
func (a *authUseCase) AuthByCredentials(ctx context.Context, login string, password string) (*usecase.AuthByCredentialsResultDto, error) {
	op := "Auth.AuthByCredentials"

	log := a.logger.With(
		slog.String("op", op),
		slog.String("login", login),
	)

	user, err := a.userSrv.FindUserByLogin(ctx, login)
	if err != nil {
		a.logger.ErrorContext(ctx, "Error from find user", utils.LogErr(err))
		return nil, err
	}

	if user == nil {
		msg := "user by provided creds not found"
		a.logger.WarnContext(ctx, msg)
		return nil, errors.New(msg)
	}

	if err := user.ComparePassword(password); err != nil {
		msg := "error from validate password"
		a.logger.WarnContext(ctx, msg, utils.LogErr(err))
		return nil, errors.New(msg)
	}

	tokens, err := a.tokenSrv.GeneratePair(user)
	if err != nil {
		log.ErrorContext(ctx, "Error from generating tokens", utils.LogErr(err))
		return nil, err
	}

	return &usecase.AuthByCredentialsResultDto{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
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

	tokens, err := a.tokenSrv.GeneratePair(user)
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

	_, err := a.tokenSrv.ValidateAccessToken(t);
  if err != nil {
		log.ErrorContext(ctx, "Error from validating token", utils.LogErr(err))
		return err
	}

	return nil
}

func (a *authUseCase) RefreshToken(ctx context.Context, token string) (*model.Token, error) {
	op := "Auth.RefreshToken"

	log := a.logger.With(slog.String("op", op))

  t := model.NewToken(token, nil)

  claims, err := a.tokenSrv.ValidateRefreshToken(t)
  if err != nil {
		log.ErrorContext(ctx, "Error from validating token", utils.LogErr(err))
		return nil, err
  }

  log.With(slog.String("uid", claims.UserId.String()))

	user, err := a.userSrv.FinUserById(ctx, claims.UserId)
	if err != nil {
		log.ErrorContext(ctx, "Error from searching user", utils.LogErr(err))
		return nil, err
	}

	res, err := a.tokenSrv.RefreshToken(user, t)
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
