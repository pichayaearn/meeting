package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/model/mocks"
	"github.com/pichayaearn/meeting/pkg/serializer"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateStatusMeeting(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name      string
		parameter serializer.UpdateMeetingReq
		mockSvc   func(m *mocks.MeetingSvc)
		isError   bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.UpdateMeetingReq{
				MeetingID: "",
				Status:    "",
			},
			isError: true,
		},
		{
			name: "failed, error from svc when update",
			parameter: serializer.UpdateMeetingReq{
				MeetingID: "a73bae08-a106-47fc-a3ee-f58736d3bd7a",
				Status:    "canceled",
			},
			mockSvc: func(m *mocks.MeetingSvc) {
				m.On("Update", mock.IsType(model.UpdateMeetingOpts{})).Return(errors.New("error from svc"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.UpdateMeetingReq{
				MeetingID: "a73bae08-a106-47fc-a3ee-f58736d3bd7a",
				Status:    "canceled",
			},
			mockSvc: func(m *mocks.MeetingSvc) {
				m.On("Update", mock.IsType(model.UpdateMeetingOpts{})).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.parameter)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodPatch, "/meeting", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", uuid.New().String())

			mockSvc := mocks.NewMeetingSvc(t)
			if tt.mockSvc != nil {
				tt.mockSvc(mockSvc)
			}

			err = UpdateStatusMeeting(UpdateStatusMeetingCfg{
				MeetingSvc: mockSvc,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(http.StatusOK, rec.Result().StatusCode)
			}
		})
	}

}
