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

type CreateMeetingOpts struct {
	Title     string
	Detail    string
	CreatedBy uuid.UUID
}
type MeetingSvc interface {
	Create(opts CreateMeetingOpts) error
	List(opts GetMeetingOpts, ctx context.Context) ([]Meeting, error)
}

type CommentSvc interface {
	List(opts GetListCommentOpts, ctx context.Context) ([]Comment, error)
}
