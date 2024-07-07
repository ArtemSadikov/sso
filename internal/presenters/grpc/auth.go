package grpc

import (
	"context"

	ssoAuth "github.com/ArtemSadikov/cinematic.back_protos/generated/go/auth"
)

type AuthServer struct {
	ssoAuth.UnimplementedAuthServiceServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s AuthServer) LoginByCredentials(
	ctx context.Context,
	in *ssoAuth.LoginByCredentialsUserRequest,
) (*ssoAuth.LoginByCredentialsUserResponse, error) {
	// TODO
	return nil, nil
}

func (s AuthServer) RegisterUser(
	ctx context.Context,
	in *ssoAuth.RegisterUserRequest,
) (*ssoAuth.RegisterUserResponse, error) {
	// TODO
	return nil, nil
}
