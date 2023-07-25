package model

import "context"

type UserRepo interface {
	Get(opts GetUserOpts, ctx context.Context) (*User, error)
	Create(user User) error
}
