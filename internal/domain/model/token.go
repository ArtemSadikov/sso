package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	UserId uuid.UUID `json:"uid"`
	jwt.RegisteredClaims
}

type Token struct {
	Value        string
	AvailableFor *time.Time
}

func BuildToken(user *User, duration time.Duration, secretKey []byte) (*Token, error) {
	availableFor := time.Now().Add(duration)

	claims := TokenClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(availableFor),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return NewToken(token, &availableFor), nil
}

func NewToken(value string, availableFor *time.Time) *Token {
	return &Token{
		Value:        value,
		AvailableFor: availableFor,
  }
}
