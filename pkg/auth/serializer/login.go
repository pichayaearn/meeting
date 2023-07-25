package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required),
	)
}

type LoginResponse struct {
	Token string `json:"token"`
}

func ToLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}
