package entity

import (
	"time"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type UserEntity struct {
	Id uuid.UUID `db:"id"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	Password string `db:"password"`
	Login    string `db:"login"`
}

func (e *UserEntity) MapToModel() *model.User {
	return &model.User{
		Id:        e.Id,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}
