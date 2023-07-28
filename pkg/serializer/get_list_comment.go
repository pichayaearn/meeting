package serializer

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"github.com/pichayaearn/meeting/pkg/model"
)

type GetListCommentReq struct {
	MeetingID string `query:"meeting_id"`
	Limit     int    `query:"limit"`
	Offset    int    `query:"offset"`
}

func (req GetListCommentReq) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.MeetingID, validation.Required, is.UUIDv4),
	)
}

type GetListCommentResponse struct {
	ID             uuid.UUID `json:"id"`
	Detail         string    `json:"detail"`
	CreatedByID    uuid.UUID `json:"created_by"`
	CreatedByEmail string    `json:"created_by_email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToGetListCommentResponse(opts model.Comment) GetListCommentResponse {
	return GetListCommentResponse{
		ID:             opts.CommentID(),
		Detail:         opts.Detail(),
		CreatedByID:    opts.CreatedByID(),
		CreatedByEmail: opts.CreatedByEmail(),
		CreatedAt:      opts.CreatedAt(),
		UpdatedAt:      opts.UpdatedAt(),
	}
}
