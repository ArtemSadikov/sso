package repository

import (
	"cinematic.com/sso/internal/domain/model"
	"cinematic.com/sso/internal/infrastructure/storage/entity"
)

func mapUserToModel(in *entity.UserEntity) *model.User {
	user := model.NewUser(in.Id, in.Login, in.Password)

	user.CreatedAt = in.CreatedAt
	user.DeletedAt = in.DeletedAt
	user.UpdatedAt = in.UpdatedAt
	user.IsDeleted = in.IsDeleted

	return user
}
