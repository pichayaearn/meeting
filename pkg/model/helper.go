package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	validator "github.com/go-ozzo/ozzo-validation"
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

type MeetingFactoryOpts struct {
	ID        uuid.UUID
	Title     string
	Detail    string
	Status    string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	DeletedAt time.Time
}

func MeetingFactory(opts MeetingFactoryOpts) (*Meeting, error) {
	meeting := Meeting{
		id:            opts.ID,
		title:         opts.Title,
		detail:        opts.Detail,
		status:        MeetingStatus(opts.Status),
		createdAt:     opts.CreatedAt,
		createdByUUID: opts.CreatedBy,
		updatedAt:     opts.UpdatedAt,
	}

	if err := meeting.Validate(validator.Field(&meeting.id, validator.Required)); err != nil {
		return nil, err
	}

	return &meeting, nil
}
