package model

import (
	"time"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type MeetingStatus string

const (
	MeetingStatusTodo       MeetingStatus = "to_do"
	MeetingStatusInProgress MeetingStatus = "in_progress"
	MeetingStatusDone       MeetingStatus = "done"
	MeetingStatusCanceled   MeetingStatus = "canceled"
)

type Meeting struct {
	id            uuid.UUID
	title         string
	detail        string
	status        MeetingStatus
	createdAt     time.Time
	createdByUUID uuid.UUID
	updatedAt     time.Time

	createdBy User
}

func (m Meeting) ID() uuid.UUID            { return m.id }
func (m Meeting) Title() string            { return m.title }
func (m Meeting) Detail() string           { return m.detail }
func (m Meeting) Status() MeetingStatus    { return m.status }
func (m Meeting) CreatedAt() time.Time     { return m.createdAt }
func (m Meeting) CreatedByUUID() uuid.UUID { return m.createdByUUID }
func (m Meeting) UpdatedAt() time.Time     { return m.updatedAt }
func (m Meeting) CreatedBy() User          { return m.createdBy }

func (m *Meeting) Validate(additionalRules ...*validator.FieldRules) error {
	rules := []*validator.FieldRules{
		validator.Field(&m.title, validator.Required),
		validator.Field(&m.detail, validator.Required),
		validator.Field(&m.status, validator.Required, validator.In(MeetingStatusTodo, MeetingStatusInProgress, MeetingStatusDone, MeetingStatusCanceled)),
		validator.Field(&m.createdAt, validator.Required),
		validator.Field(&m.createdByUUID, validator.Required),
	}

	if additionalRules != nil {
		rules = append(rules, additionalRules...)
	}

	if err := validator.ValidateStruct(m, rules...); err != nil {
		return err
	}

	return nil
}

func (m *Meeting) SetCreatedBy(user User) error {
	m.createdBy = user
	return m.Validate()
}

func NewMeeting(opts CreateMeetingOpts) (*Meeting, error) {
	now := time.Now()
	meeting := Meeting{
		title:         opts.Title,
		detail:        opts.Detail,
		status:        MeetingStatusTodo,
		createdAt:     now,
		createdByUUID: opts.CreatedBy,
		updatedAt:     now,
	}

	if err := meeting.Validate(); err != nil {
		return nil, err
	}

	return &meeting, nil
}
