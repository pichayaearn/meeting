package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	model "github.com/pichayaearn/meeting/pkg/model"
)

type CreateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateUserReq) ValidateRequest() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required),
	)
}

type CreateUserResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
}

func ToCreateUserResponse(user model.User) CreateUserResponse {
	return CreateUserResponse{
		UserID: user.UserID(),
		Email:  user.Email(),
	}
}
