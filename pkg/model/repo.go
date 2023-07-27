package model

import (
	"context"

	"github.com/google/uuid"
)

type UserRepo interface {
	Get(opts GetUserOpts, ctx context.Context) (*User, error)
	Create(user User) error
}

type GetMeetingOpts struct {
	ID        uuid.UUID
	Status    string
	CreatedBy uuid.UUID
	Limit     int
	Offset    int
}

type MeetingRepo interface {
	Get(opts GetMeetingOpts, ctx context.Context) (*Meeting, error)
	List(opts GetMeetingOpts, ctx context.Context) ([]Meeting, error)
	Create(meeting Meeting) error
}
