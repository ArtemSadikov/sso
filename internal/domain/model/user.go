package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id uuid.UUID

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	password string
	Login    string

	Contacts []*UserContact
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) AddContact(contact *UserContact) {
	u.Contacts = append(u.Contacts, contact)
}

func (u *User) ComparePassword(pswd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(pswd))
}

func NewUser(id uuid.UUID, login, password string, contacts ...*UserContact) *User {
	return &User{
		Id:        id,
		Login:     login,
		password:  password,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}
}
