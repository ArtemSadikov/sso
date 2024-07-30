package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	password string

	Contacts []*UserContact
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) AddContact(contact *UserContact) {
	u.Contacts = append(u.Contacts, contact)
}

func NewUserWithoutId(contacts ...*UserContact) *User {
	return NewUser(uuid.New(), contacts...)
}

func NewUser(id uuid.UUID, contacts ...*UserContact) *User {
	return &User{
		Id:        id,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}
}
