package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID

	Login    string
	Password string
}

func NewUserWithId(id uuid.UUID, login string) *User {
	return &User{
		Id: id,
		Login: login,
	}
}

func NewUserWithPassword(login string, password string) *User {
	user := NewUser(login)
	user.Password = password
	return user
}

func NewUser(login string) *User {
	return NewUserWithId(uuid.New(), login)
}
