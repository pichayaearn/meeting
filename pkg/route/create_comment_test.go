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

func TestCreateComment(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name      string
		parameter serializer.CreateCommentReq
		mockSvc   func(m *mocks.CommentSvc)
		isError   bool
	}{
		{
			name: "failed, invalid parameter",
			parameter: serializer.CreateCommentReq{
				Detail:    "",
				MeetingID: "",
				CreatedBy: uuid.Nil,
			},
			isError: true,
		},
		{
			name: "failed, error from create comment svc",
			parameter: serializer.CreateCommentReq{
				Detail:    "Detail",
				CreatedBy: uuid.New(),
				MeetingID: uuid.New().String(),
			},
			mockSvc: func(m *mocks.CommentSvc) {
				m.On("Create", mock.IsType(model.CreateCommentOpts{})).Return(errors.New("error when create"))
			},
			isError: true,
		},
		{
			name: "success",
			parameter: serializer.CreateCommentReq{
				Detail:    "Detail",
				CreatedBy: uuid.New(),
				MeetingID: uuid.New().String(),
			},
			mockSvc: func(m *mocks.CommentSvc) {
				m.On("Create", mock.IsType(model.CreateCommentOpts{})).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.parameter)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodPost, "/comment", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)
			c.Set("ActionBy", tt.parameter.CreatedBy.String())

			mockSvc := mocks.NewCommentSvc(t)
			if tt.mockSvc != nil {
				tt.mockSvc(mockSvc)
			}

			err = CreateComment(CreateCommentCfg{
				CommentSvc: mockSvc,
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
