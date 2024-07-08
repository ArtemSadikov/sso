package container

import (
	"fmt"
	"log/slog"
	"net"

	"cinematic.com/sso/internal/domain/service"
	"cinematic.com/sso/internal/infrastructure/config"
	"cinematic.com/sso/internal/infrastructure/logger"
	grpcServers "cinematic.com/sso/internal/presenters/grpc"
	"cinematic.com/sso/internal/usecase"
	uc "cinematic.com/sso/internal/usecase"
	ssoAuthApi "github.com/ArtemSadikov/cinematic.back_protos/generated/go/auth"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Container struct {
	container *dig.Container
}

func New(opts ...dig.Option) (*Container, error) {
	container := dig.New(opts...)

	if err := container.Provide(func() *config.Config {
		return config.MustLoad()
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(cfg *config.Config) *slog.Logger {
		return logger.SetupLogger(cfg.Env)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(logger *slog.Logger) service.UserService {
		return service.NewUserService(logger)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(logger *slog.Logger, userSrv service.UserService) usecase.AuthUseCase {
		return uc.NewAuthUseCase(logger, userSrv)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(authUc usecase.UseCases) grpcServers.AuthServer {
		return *grpcServers.NewAuthServer(authUc)
	}); err != nil {
		return nil, err
	}

	return &Container{container: container}, nil
}

func (c *Container) Run() error {
	if err := c.container.Invoke(func(
		cfg *config.Config,
		logger *slog.Logger,
		authServer grpcServers.AuthServer,
	) error {
		srvErr := make(chan error)
		s := grpc.NewServer()

		ssoAuthApi.RegisterAuthServiceServer(s, authServer)
		reflection.Register(s)

		go func() {
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
			if err != nil {
				srvErr <- err
				return
			}

			logger.Info("Start serving grpc...", slog.Int("port", cfg.GRPC.Port))
			if err := s.Serve(listener); err != nil {
				srvErr <- err
				return
			}
		}()

		return <-srvErr
	}); err != nil {
		return err
	}

	return nil
}
