package repository

import (
	"context"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/infrastructure/storage/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	logger  *slog.Logger
	db *sqlx.DB
}

// FindUsersByIds implements UserRepository.
func (u *userRepo) FindUsersByIds(ctx context.Context, ids ...string) ([]*model.User, error) {
	return nil, nil
}

// RemoveUsers implements UserRepository.
func (u *userRepo) RemoveUsers(ctx context.Context, users ...*model.User) error {
	return nil
}

func (u *userRepo) CreateUser(ctx context.Context, login string, password string) (*uuid.UUID, error) {
	user := &entity.UserEntity{}

	q, err := u.db.PrepareContext(ctx, "INSERT INTO sso.users(login, password) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := q.ExecContext(ctx, login, password)
	if err != nil {
		return nil, err
	}

	return uuid.MustParse(res.LastInsertId()), nil
}

func (u *userRepo) UpdateUser(ctx context.Context, login string) (*model.User, error) {
	return nil, nil
}

func NewUserRepository(
	logger *slog.Logger,
	db *sqlx.DB,
) *userRepo {
	return &userRepo{logger, db}
}
