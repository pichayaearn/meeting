package route

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/model/mocks"
	"github.com/pichayaearn/meeting/pkg/serializer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetListMeeting(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name      string
		parameter serializer.GetListMeetingReq
		mockSvc   func(m *mocks.MeetingSvc)
		isError   bool
	}{
		{
			name: "failed, validate parameter failed",
			parameter: serializer.GetListMeetingReq{
				ID: "id not uuid",
			},

			isError: true,
		},
		{
			name:      "failed, error from get list meeting svc",
			parameter: serializer.GetListMeetingReq{},
			mockSvc: func(m *mocks.MeetingSvc) {
				m.On("List", mock.IsType(model.GetMeetingOpts{}), context.Background()).Return(nil, errors.New("get list meeting error"))
			},

			isError: true,
		},
		{
			name:      "success",
			parameter: serializer.GetListMeetingReq{},
			mockSvc: func(m *mocks.MeetingSvc) {
				meeting, err := model.NewMeeting(model.CreateMeetingOpts{
					Title:     "test",
					Detail:    "test",
					CreatedBy: uuid.New(),
				})
				r.NoError(err)
				lists := []model.Meeting{}
				lists = append(lists, *meeting)
				m.On("List", mock.IsType(model.GetMeetingOpts{}), context.Background()).Return(lists, nil)
			},

			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/meetings", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			// Set query parameters
			q := req.URL.Query()
			q.Set("id", tt.parameter.ID)
			q.Set("status", tt.parameter.Status)
			q.Set("created_by", tt.parameter.CreatedBy)
			q.Set("limit", fmt.Sprintf("%d", tt.parameter.Limit))
			q.Set("offset", fmt.Sprintf("%d", tt.parameter.Offset))
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			mockMeetingSvc := mocks.NewMeetingSvc(t)

			if tt.mockSvc != nil {
				tt.mockSvc(mockMeetingSvc)
			}

			err := GetListMeeting(GetListMeetingCfg{
				MeetingSvc: mockMeetingSvc,
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
