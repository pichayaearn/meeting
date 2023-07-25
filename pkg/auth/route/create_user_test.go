package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/auth/serializer"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/model/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		name    string
		body    serializer.CreateUserReq
		mockSvc func(m *mocks.UserSvc)
		isError bool
	}{
		{
			name: "failed, invalid paremater",
			body: serializer.CreateUserReq{
				Email:    "",
				Password: "",
			},
			isError: true,
		},
		{
			name: "failed, error from svc",
			body: serializer.CreateUserReq{
				Email:    "test1@gmail.com",
				Password: "test1",
			},
			mockSvc: func(m *mocks.UserSvc) {
				m.On("CreateUser", model.CreateUser{
					Email:    "test1@gmail.com",
					Password: "test1",
				}).Return(errors.New("error from user svc"))
			},
			isError: true,
		},
		{
			name: "success",
			body: serializer.CreateUserReq{
				Email:    "test1@gmail.com",
				Password: "test1",
			},
			mockSvc: func(m *mocks.UserSvc) {

				m.On("CreateUser", model.CreateUser{
					Email:    "test1@gmail.com",
					Password: "test1",
				}).Return(nil)
			},
			isError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, err := json.Marshal(tt.body)
			r.NoError(err)
			req := httptest.NewRequest(http.MethodGet, "/login", strings.NewReader(string(reqBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e := echo.New()
			c := e.NewContext(req, rec)

			mockSvc := mocks.NewUserSvc(t)
			if tt.mockSvc != nil {
				tt.mockSvc(mockSvc)
			}

			err = CreateUser(CreateUserCfg{
				UserSvc: mockSvc,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(http.StatusCreated, rec.Code)
			}
		})

	}
}
