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

type MeetingCommentFactoryOpts struct {
	ID        uuid.UUID
	MeetingID uuid.UUID
	CommentID uuid.UUID
	Status    string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
}

func MeetingCommentFactory(opts MeetingCommentFactoryOpts) (*MeetingComment, error) {
	meetingComment := MeetingComment{
		id:          opts.ID,
		meetingID:   opts.MeetingID,
		commentID:   opts.CommentID,
		status:      StatusComment(opts.Status),
		createdByID: opts.CreatedBy,
		createdAt:   opts.CreatedAt,
		updatedAt:   opts.UpdatedAt,
	}
	if err := meetingComment.Validate(); err != nil {
		return nil, err
	}

	return &meetingComment, nil
}

type CommentDetailFactoryOpts struct {
	ID     uuid.UUID
	Detail string
}

func CommentDetailFactory(opts CommentDetailFactoryOpts) (*CommentDetail, error) {
	comment := CommentDetail{
		id:     opts.ID,
		detail: opts.Detail,
	}

	if err := comment.Validate(validator.Field(&comment.id, validator.Required)); err != nil {
		return nil, err
	}

	return &comment, nil
}
