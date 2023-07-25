package route

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/auth/model/mocks"
	"github.com/pichayaearn/meeting/pkg/auth/serializer"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	r := require.New(t)

	tests := []struct {
		name    string
		body    serializer.LoginReq
		mockSvc func(m *mocks.AuthSvc)
		isError bool
	}{
		{
			name: "failed, invalid paremater",
			body: serializer.LoginReq{
				Email:    "",
				Password: "",
			},
			isError: true,
		},
		{
			name: "failed, error from svc",
			body: serializer.LoginReq{
				Email:    "test1@gmail.com",
				Password: "test1",
			},
			mockSvc: func(m *mocks.AuthSvc) {
				m.On("Login", "test1@gmail.com", "test1").Return("", errors.New("error from auth svc"))
			},
			isError: true,
		},
		{
			name: "success",
			body: serializer.LoginReq{
				Email:    "test1@gmail.com",
				Password: "test1",
			},
			mockSvc: func(m *mocks.AuthSvc) {
				m.On("Login", "test1@gmail.com", "test1").Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJkZDlhMTI2ZS05ZWI2LTQwMzAtYWFlMy00YTBmMjk5MjI5ZjIiLCJleHAiOjE2OTAyOTUzMTV9.xyMlFpJ2PdB04cz1ztMJHktBEjvC7a5mlkBh-Aeg9gA", nil)
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

			mockSvc := mocks.NewAuthSvc(t)
			if tt.mockSvc != nil {
				tt.mockSvc(mockSvc)
			}

			err = Login(LoginCfg{
				AuthSvc: mockSvc,
			})(c)

			if tt.isError {
				r.Error(err)
			} else {
				r.NoError(err)
				r.NotEmpty(rec.Body)
				r.Equal(http.StatusOK, rec.Code)
			}
		})

	}
}
