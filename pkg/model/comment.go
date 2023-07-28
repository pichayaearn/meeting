package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type StatusComment string

const (
	StatusCommentActive   StatusComment = "active"
	StatusCommentUnActive StatusComment = "unactive"
)

type MeetingComment struct {
	id             uuid.UUID
	meetingID      uuid.UUID
	commentID      uuid.UUID
	status         StatusComment
	createdByID    uuid.UUID
	createdByEmail string
	createdAt      time.Time
	updatedAt      time.Time
}

func (mc MeetingComment) ID() uuid.UUID          { return mc.id }
func (mc MeetingComment) MeetingID() uuid.UUID   { return mc.meetingID }
func (mc MeetingComment) CommentID() uuid.UUID   { return mc.commentID }
func (mc MeetingComment) Status() StatusComment  { return mc.status }
func (mc MeetingComment) CreatedByID() uuid.UUID { return mc.createdByID }
func (mc MeetingComment) CreatedByEmail() string { return mc.createdByEmail }
func (mc MeetingComment) CreatedAt() time.Time   { return mc.createdAt }
func (mc MeetingComment) UpdatedAt() time.Time   { return mc.updatedAt }

func (mc *MeetingComment) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&mc.id, validator.Required),
		validator.Field(&mc.meetingID, validator.Required),
		validator.Field(&mc.commentID, validator.Required),
		validator.Field(&mc.status, validator.Required, validator.In(StatusCommentActive, StatusCommentUnActive)),
		validator.Field(&mc.createdAt, validator.Required),
		validator.Field(&mc.createdByID, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(mc, rules...); err != nil {
		return err
	}

	return nil
}

func (mc *MeetingComment) SetCreatedBy(email string) error {
	mc.createdByEmail = email
	if err := mc.Validate(validator.Field(&mc.createdByEmail, validator.Required)); err != nil {
		return err
	}
	return nil
}

type CommentDetail struct {
	id     uuid.UUID
	detail string
}

func (cd CommentDetail) ID() uuid.UUID  { return cd.id }
func (cd CommentDetail) Detail() string { return cd.detail }

func (cd *CommentDetail) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&cd.detail, validator.Required),
	}
	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}
	if err := validator.ValidateStruct(cd, rules...); err != nil {
		return err
	}

	return nil
}

// type Comment struct {
// 	ID            uuid.UUID
// 	Status        StatusComment
// 	Detail        string
// 	CreatedByID   uuid.UUID
// 	CretedByEmail string
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// }

type Comment struct {
	MeetingComment
	CommentDetail
}

func (c *Comment) SetMeetingComment(data MeetingComment) {
	c.MeetingComment = data
}
func (c *Comment) SetCommentDetail(data CommentDetail) {
	c.CommentDetail = data
}
