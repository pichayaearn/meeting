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
	Update(meeting Meeting) error
}

type GetListCommentOpts struct {
	MeetingID uuid.UUID
	Limit     int
	Offset    int
}
type CommentRepo interface {
	ListCommentID(opts GetListCommentOpts, ctx context.Context) ([]MeetingComment, error)
	CommentDetail(id uuid.UUID, ctx context.Context) (*CommentDetail, error)
	CreateCommentDetail(comment CommentDetail) (*CommentDetail, error)
	CreateCommentMeeting(meetingComment MeetingComment) error
}
