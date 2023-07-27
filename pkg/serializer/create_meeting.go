package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type CreateMeetingReq struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`

	CreatedBy uuid.UUID `json:"-"`
}

func NewCreateMeetingReq(createdBy uuid.UUID) *CreateMeetingReq {
	return &CreateMeetingReq{
		CreatedBy: createdBy,
	}
}

func (req CreateMeetingReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CreatedBy, validation.Required, is.UUIDv4),
		validation.Field(&req.Title, validation.Required),
		validation.Field(&req.Detail, validation.Required),
	)
}
