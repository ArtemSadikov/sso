package repository

import (
	"context"
	"database/sql"
	"log/slog"

	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/infrastructure/storage/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	logger *slog.Logger
	db     *sqlx.DB
}

func (u *userRepo) FindUserByLogin(ctx context.Context, login string) (*model.User, error) {
	res := entity.UserEntity{}

	if err := u.db.GetContext(ctx, &res, `SELECT * FROM sso.users WHERE login=$1`, login); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return mapUserToModel(&res), nil
}

// FindUsersByIds implements UserRepository.
func (u *userRepo) FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	res := entity.UserEntity{}

	if err := u.db.GetContext(ctx, &res, "SELECT * FROM sso.users WHERE id = $1 AND is_deleted", id.String()); err != nil {
		return nil, err
	}

	return mapUserToModel(&res), nil
}

// RemoveUsers implements UserRepository.
func (u *userRepo) RemoveUsers(ctx context.Context, users ...*model.User) error {
	return nil
}

func (u *userRepo) CreateUser(ctx context.Context, login, password string, contacts ...*model.UserContact) (*model.User, error) {
	user := &entity.UserEntity{}

	tx, err := u.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	uq, err := tx.PreparexContext(ctx, `
		INSERT INTO sso.users(login, password)
		VALUES ($1, $2) RETURNING *
	`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := uq.QueryRowxContext(ctx, login, password).StructScan(user); err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(contacts) > 0 {
		c := make([]*entity.UserContact, len(contacts))
		for _, contact := range contacts {
			c = append(c, entity.NewUserContactFromModel(contact))
		}

		q := "INSERT INTO sso.user_contacts(id, _type, _value, user_id) VALUES (:id, :_type, :_value, :user_id)"
		if _, err := tx.NamedExecContext(ctx, q, c); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user.MapToModel(), nil
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
