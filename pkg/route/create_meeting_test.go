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

func TestCreateMeeting(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name      string
		parameter serializer.CreateMeetingReq
		mockSvc   func(m *mocks.MeetingSvc)
		isError   bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.CreateMeetingReq{
				Title:     "",
				Detail:    "",
				CreatedBy: uuid.Nil,
			},
			isError: true,
		},
		{
			name: "failed, error from create meeting svc",
			parameter: serializer.CreateMeetingReq{
				Title:     "Test",
				Detail:    "Detail",
				CreatedBy: uuid.New(),
			},
			mockSvc: func(m *mocks.MeetingSvc) {
				m.On("Create", mock.IsType(model.CreateMeetingOpts{})).Return(errors.New("error when create"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.CreateMeetingReq{
				Title:     "Test",
				Detail:    "Detail",
				CreatedBy: uuid.New(),
			},
			mockSvc: func(m *mocks.MeetingSvc) {
				m.On("Create", mock.IsType(model.CreateMeetingOpts{})).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.parameter)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodPost, "/meeting", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", tt.parameter.CreatedBy.String())

			mockMeetingSvc := mocks.NewMeetingSvc(t)
			if tt.mockSvc != nil {
				tt.mockSvc(mockMeetingSvc)
			}

			err = CreateMeeting(CreateMeetingCfg{
				MeetingSvc: mockMeetingSvc,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(http.StatusCreated, rec.Result().StatusCode)
			}

		})
	}
}
