package auth

import (
	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/presenters/grpc"
	"github.com/ArtemSadikov/cinematic.back_protos/generated/go/sso"
)

func mapRegisterResponse(user *model.User, accessToken *model.Token, refreshToken *model.Token) *sso.RegisterUserResponse {
	u := grpc.MapUserFromModel(user)
	at := grpc.MapTokenFromModel(accessToken)
	rt := grpc.MapTokenFromModel(refreshToken)

	return &sso.RegisterUserResponse{User: u, AccessToken: at, RefreshToken: rt}
}
