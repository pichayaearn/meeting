package serializer

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/pichayaearn/meeting/pkg/model"
)

type GetListMeetingReq struct {
	ID        string `query:"id"`
	Status    string `query:"status"`
	CreatedBy string `query:"created_by"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

func (req GetListMeetingReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.CreatedBy, is.UUIDv4),
		validation.Field(&req.ID, is.UUIDv4),
		validation.Field(&req.Status, validation.In(string(model.MeetingStatusTodo), string(model.MeetingStatusInProgress), string(model.MeetingStatusDone), string(model.MeetingStatusCanceled))),
	)
}

type GetListMeetingResponse struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Detail         string    `json:"detail"`
	Status         string    `json:"status"`
	CreatedByID    uuid.UUID `json:"created_by_id"`
	CreatedByEmail string    `json:"created_by_email"`
	CreatedAt      time.Time `json:"created_at"`
}

func ToGetListMeetingResponse(meeting model.Meeting) GetListMeetingResponse {
	return GetListMeetingResponse{
		ID:             meeting.ID(),
		Title:          meeting.Title(),
		Detail:         meeting.Detail(),
		Status:         string(meeting.Status()),
		CreatedByID:    meeting.CreatedByUUID(),
		CreatedByEmail: meeting.CreatedBy().Email(),
		CreatedAt:      meeting.CreatedAt(),
	}
}
