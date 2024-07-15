package entity

import (
	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type UserEntity struct {
	Id string `db:"id"`

	Login    string `db:"login"`
	Password string
}

func MapToModel(entity *UserEntity) *model.User {
	return model.NewUserWithId(uuid.MustParse(entity.Id), entity.Login)
}
