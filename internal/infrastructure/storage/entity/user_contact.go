package entity

import (
	"time"

	"cinematic.com/sso/internal/domain/model"
	"github.com/google/uuid"
)

type UserContact struct {
	Id uuid.UUID `db:"id"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`

	Type  string `db:"_type"`
	Value string `db:"_value"`

	UserId uuid.UUID `db:"user_id"`
}

func NewUserContactFromModel(m *model.UserContact) *UserContact {
	return &UserContact{
		Id:     m.Id,
		Type:   m.Type.String(),
		Value:  m.Value,
		UserId: m.UserId,
	}
}
