package route

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/model/mocks"
	"github.com/pichayaearn/meeting/pkg/serializer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetListComment(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name      string
		parameter serializer.GetListCommentReq
		mockSvc   func(m *mocks.CommentSvc)
		isError   bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.GetListCommentReq{
				MeetingID: "",
			},
			isError: true,
		},
		{
			name: "failed, error from get comment svc",
			parameter: serializer.GetListCommentReq{
				MeetingID: "a73bae08-a106-47fc-a3ee-f58736d3bd7a",
			},
			mockSvc: func(m *mocks.CommentSvc) {
				m.On("List", mock.IsType(model.GetListCommentOpts{}), context.Background()).Return(nil, errors.New("error from svc"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.GetListCommentReq{
				MeetingID: "a73bae08-a106-47fc-a3ee-f58736d3bd7a",
			},
			mockSvc: func(m *mocks.CommentSvc) {

				m.On("List", mock.IsType(model.GetListCommentOpts{}), context.Background()).Return(nil, nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/comments", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			// Set query parameters
			q := req.URL.Query()
			q.Set("meeting_id", tt.parameter.MeetingID)
			q.Set("limit", fmt.Sprintf("%d", tt.parameter.Limit))
			q.Set("offset", fmt.Sprintf("%d", tt.parameter.Offset))
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			mockSvc := mocks.NewCommentSvc(t)
			if tt.mockSvc != nil {
				tt.mockSvc(mockSvc)
			}

			err := GetListComment(GetListCommentCfg{
				CommentSvc: mockSvc,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.NotEmpty(rec.Body)
				r.Equal(http.StatusOK, rec.Result().StatusCode)
			}
		})
	}
}
