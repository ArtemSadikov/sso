package service

import (
	"cinematic.com/sso/internal/domain/model"
)

type TokenPair struct {
	AccessToken  *model.Token
	RefreshToken *model.Token
}
