package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type CreateCommentReq struct {
	Detail    string `json:"detail"`
	MeetingID string `json:"meeting_id"`

	CreatedBy uuid.UUID `json:"-"`
}

func NewCreateCommentReq(createdBy uuid.UUID) *CreateCommentReq {
	return &CreateCommentReq{
		CreatedBy: createdBy,
	}
}

func (req CreateCommentReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CreatedBy, validation.Required, is.UUIDv4),
		validation.Field(&req.MeetingID, validation.Required, is.UUIDv4),
		validation.Field(&req.Detail, validation.Required),
	)
}
