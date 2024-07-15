package grpc

import (
	"context"

	"cinematic.com/sso/internal/usecase"
	ssoAuth "github.com/ArtemSadikov/cinematic.back_protos/generated/go/auth"
)

type AuthServer struct {
	uc usecase.AuthUseCase

	ssoAuth.UnimplementedAuthServiceServer
}

func NewAuthServer(
	uc usecase.AuthUseCase,
) *AuthServer {
	return &AuthServer{uc: uc}
}

func (s AuthServer) LoginByCredentials(
	ctx context.Context,
	in *ssoAuth.LoginByCredentialsUserRequest,
) (*ssoAuth.LoginByCredentialsUserResponse, error) {
	s.uc.AuthByCredentials(ctx, in.Login, in.Password)
	// TODO
	return nil, nil
}

func (s AuthServer) RegisterUser(
	ctx context.Context,
	in *ssoAuth.RegisterUserRequest,
) (*ssoAuth.RegisterUserResponse, error) {
	res, err := s.uc.RegisterUser(ctx, in.Login, in.Password)
	if err != nil {
		return nil, err
	}
	// TODO
	return &ssoAuth.RegisterUserResponse{
		Id:    res.Id.String(),
		Login: res.Login,
	}, nil
}
