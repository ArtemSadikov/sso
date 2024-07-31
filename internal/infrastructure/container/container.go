package container

import (
	"fmt"
	"log/slog"
	"net"

	"cinematic.com/sso/internal/domain/service"
	"cinematic.com/sso/internal/infrastructure/config"
	"cinematic.com/sso/internal/infrastructure/logger"
	"cinematic.com/sso/internal/infrastructure/storage/postgresql"
	"cinematic.com/sso/internal/infrastructure/storage/repository"
	authSrv "cinematic.com/sso/internal/presenters/grpc/auth"
	authUc "cinematic.com/sso/internal/usecase/auth"
	"cinematic.com/sso/internal/usecase"
	userRepo "cinematic.com/sso/internal/infrastructure/storage/repository/user"
	"github.com/ArtemSadikov/cinematic.back_protos/generated/go/sso"
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

	if err := container.Provide(func(cfg *config.Config, logger *slog.Logger) (*postgresql.Storage, error) {
		connection := postgresql.NewPostgreSQLStorage(cfg.Storage, logger)
		err := connection.Connect()
		return connection, err
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(cfg *config.Config) service.TokenService {
		return service.NewTokenService(cfg)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(logger *slog.Logger, storage *postgresql.Storage) repository.UserRepository {
		return userRepo.NewUserRepository(logger, storage.DB)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(logger *slog.Logger, repo repository.UserRepository) service.UserService {
		return service.NewUserService(logger, repo)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(logger *slog.Logger, userSrv service.UserService, tokenSrv service.TokenService) usecase.AuthUseCase {
		return authUc.NewAuthUseCase(logger, userSrv, tokenSrv)
	}); err != nil {
		return nil, err
	}

	if err := container.Provide(func(authUc usecase.AuthUseCase) authSrv.AuthServer {
		return *authSrv.NewAuthServer(authUc)
	}); err != nil {
		return nil, err
	}

	return &Container{container: container}, nil
}

func (c *Container) Run() error {
	if err := c.container.Invoke(func(
		cfg *config.Config,
		logger *slog.Logger,
		authServer authSrv.AuthServer,
	) error {
		srvErr := make(chan error)
		s := grpc.NewServer()

		sso.RegisterAuthServiceServer(s, authServer)
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
