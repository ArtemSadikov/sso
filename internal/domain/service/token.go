package service

import (
	"time"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
	cfg *config.Config
}

func (s *tokenService) GeneratePair(user *model.User) (*TokenPair, error) {
	acT, err := model.BuildToken(user, s.cfg.TokenTTL, []byte(s.cfg.AccessTokenSecret))
	if err != nil {
		return nil, err
	}

	rfT, err := model.BuildToken(user, 24*time.Hour, []byte(s.cfg.RefreshTokenSecret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  acT,
		RefreshToken: rfT,
	}, nil
}

func (s *tokenService) RefreshToken(user *model.User, token *model.Token) (*model.Token, error) {
	if _, err := s.ValidateRefreshToken(token); err != nil {
		return nil, err
	}

	res, err := model.BuildToken(user, s.cfg.TokenTTL, []byte(s.cfg.AccessTokenSecret))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *tokenService) ValidateRefreshToken(token *model.Token) (*model.TokenClaims, error) {
	return s.validate(token, s.cfg.RefreshTokenSecret)
}

func (s *tokenService) ValidateAccessToken(token *model.Token) (*model.TokenClaims, error) {
	return s.validate(token, s.cfg.AccessTokenSecret)
}

func (s *tokenService) validate(token *model.Token, secret string) (*model.TokenClaims, error) {
	res, err := jwt.ParseWithClaims(token.Value, &model.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return res.Claims.(*model.TokenClaims), nil
}

func NewTokenService(cfg *config.Config) *tokenService {
	return &tokenService{cfg}
}
