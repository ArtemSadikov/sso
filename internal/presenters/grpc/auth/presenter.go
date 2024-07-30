package auth

import (
	"context"

	"cinematic.com/sso/internal/presenters/grpc"
	"cinematic.com/sso/internal/usecase"
	"github.com/ArtemSadikov/cinematic.back_protos/generated/go/sso"
	"github.com/google/uuid"
)

type AuthServer struct {
	uc usecase.AuthUseCase

	sso.UnimplementedAuthServiceServer
}

func NewAuthServer(
	uc usecase.AuthUseCase,
) *AuthServer {
	return &AuthServer{uc: uc}
}

func (s AuthServer) LoginByCredentials(
	ctx context.Context,
	in *sso.LoginByCredsReq,
) (*sso.LoginByCredsRes, error) {
	s.uc.AuthByCredentials(ctx, in.Credentials.GetLogin(), in.Credentials.GetPassword())
	// TODO
	return nil, nil
}

func (s AuthServer) ValidateToken(
	ctx context.Context,
	in *sso.ValidateTokenReq,
) (*sso.ValidateTokenRes, error) {
	if err := s.uc.ValidateToken(ctx, in.GetToken()); err != nil {
		return nil, err
	}

	return &sso.ValidateTokenRes{Ok: true}, nil
}

func (s AuthServer) RegisterUser(
	ctx context.Context,
	in *sso.RegisterUserRequest,
) (*sso.RegisterUserResponse, error) {
	res, err := s.uc.RegisterUser(ctx, in.Credentials.GetLogin(), in.Credentials.GetPassword())
	if err != nil {
		return nil, err
	}
	// TODO
	return mapRegisterResponse(res.User, res.AccessToken, res.RefreshToken), nil
}

func (s AuthServer) RefreshToken(
	ctx context.Context,
	in *sso.RefreshTokenReq,
) (*sso.RefreshTokenRes, error) {
	userId, err := uuid.Parse(in.GetUserId())
	if err != nil {
		return nil, err
	}
	res, err := s.uc.RefreshToken(ctx, in.GetRefreshToken(), userId)
	if err != nil {
		return nil, err
	}

	return &sso.RefreshTokenRes{
		AccessToken: grpc.MapTokenFromModel(res),
	}, nil
}
