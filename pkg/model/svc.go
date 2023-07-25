package model

import (
	"context"

	"github.com/google/uuid"
)

type GetUserOpts struct {
	UserID uuid.UUID
	Email  string
	Status UserStatus
}

type CreateUser struct {
	Email    string
	Password string
}

type UserSvc interface {
	GetUser(opts GetUserOpts, ctx context.Context) (*User, error)
	CreateUser(opts CreateUser) error
}
