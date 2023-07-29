package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pichayaearn/meeting/pkg/model"
)

type UpdateMeetingReq struct {
	MeetingID string `json:"meeting_id"`
	Status    string `json:"status"`
}

func (req UpdateMeetingReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.MeetingID, validation.Required, is.UUIDv4),
		validation.Field(&req.Status, validation.Required, validation.In(string(model.MeetingStatusTodo), string(model.MeetingStatusInProgress), string(model.MeetingStatusDone), string(model.MeetingStatusCanceled))),
	)
}
