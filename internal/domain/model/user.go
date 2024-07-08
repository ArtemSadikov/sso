package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID

	Login string
}

func NewUserWithId(id uuid.UUID, login string) *User {
	return &User{id, login}
}

func NewUser(login string) *User {
	return &User{
		Id:    uuid.New(),
		Login: login,
	}
}
