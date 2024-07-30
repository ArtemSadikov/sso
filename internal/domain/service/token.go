package service

import (
	"context"
	"errors"
	"time"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
  cfg *config.Config
}

func (s *tokenService) GeneratePair(ctx context.Context, user *model.User) (*TokenPair, error) {
  acT, err := model.BuildToken(user, s.cfg.TokenTTL, []byte(s.cfg.TokenSecret))
	if err != nil {
		return nil, err
	}

  rfT, err := model.BuildToken(user, 24 * time.Hour, []byte(s.cfg.TokenSecret))
	if err != nil {
		return nil, err
	}

  return &TokenPair{
    AccessToken: acT,
    RefreshToken: rfT ,
  }, nil
}

func (s *tokenService) RefreshToken(ctx context.Context, user *model.User, token *model.Token) (*model.Token, error) {
  if err := s.Validate(ctx, token); err != nil {
    return nil, err
  }

  res, err := model.BuildToken(user, s.cfg.TokenTTL, []byte(s.cfg.TokenSecret))
	if err != nil {
		return nil, err
	}

  return res, nil
}

func (s *tokenService) Validate(ctx context.Context, token *model.Token) error {
  res, err := jwt.Parse(token.Value, func(t *jwt.Token) (interface{}, error) {
    return []byte(s.cfg.TokenSecret), nil
  })

  if !res.Valid {
    return errors.New("invalid token")
  }

  return err
}

func NewTokenService(
  cfg *config.Config,
) *tokenService {
  return &tokenService{cfg}
}
