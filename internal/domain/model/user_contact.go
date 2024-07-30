package model

import (
	"time"

	"github.com/google/uuid"
)

type UserContactType string

func (t UserContactType) String() string {
	return string(t)
}

const (
	Login UserContactType = "LOGIN"
)

type UserContact struct {
	Id uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Value string
	Type  UserContactType

	UserId uuid.UUID
}

func NewUserContactWithoutId(value string, _type UserContactType, userId uuid.UUID) *UserContact {
	return NewUserContact(uuid.New(), value, _type, userId)
}

func NewUserContact(id uuid.UUID, value string, _type UserContactType, userId uuid.UUID) *UserContact {
	return &UserContact{
		Id:     id,
		Value:  value,
		Type:   _type,
		UserId: userId,
	}
}
