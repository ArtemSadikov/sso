package usecase

import "cinematic.com/sso/internal/domain/model"

type RegisterResultDto struct {
	User         *model.User
	AccessToken  *model.Token
	RefreshToken *model.Token
}
