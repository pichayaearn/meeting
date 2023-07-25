package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type UserFactoryOpts struct {
	UserID    uuid.UUID
	Email     string
	Password  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func UserFactory(opts UserFactoryOpts) (*User, error) {
	user := User{
		userID:    opts.UserID,
		email:     opts.Email,
		password:  opts.Password,
		status:    UserStatus(opts.Status),
		createdAt: opts.CreatedAt,
		updatedAt: opts.UpdatedAt,
	}

	if err := user.Validate(validation.Field(&user.userID, validation.Required, is.UUID)); err != nil {
		return nil, err
	}

	return &user, nil
}
