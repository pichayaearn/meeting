package auth

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	authModel "github.com/pichayaearn/meeting/pkg/auth/model"
	model "github.com/pichayaearn/meeting/pkg/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthSvc struct {
	userSvc   model.UserSvc
	secretKey string
}

type NewAuthSvcCfg struct {
	UserSvc   model.UserSvc
	SecretKey string
}

type MyCustomClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func NewAuthSvc(cfg NewAuthSvcCfg) authModel.AuthSvc {
	return &AuthSvc{
		userSvc:   cfg.UserSvc,
		secretKey: cfg.SecretKey,
	}
}

func (aSvc AuthSvc) Login(email, password string) (string, error) {
	ctx := context.Background()
	userExist, err := aSvc.userSvc.GetUser(model.GetUserOpts{
		Email:  email,
		Status: model.UserStatusActived,
	}, ctx)
	if err != nil {
		return "", err
	}

	if userExist == nil {
		return "", fmt.Errorf("email %s not found", email)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userExist.Password()), []byte(password)); err != nil {
		return "", err
	}

	return aSvc.createToken(userExist.UserID())
}

func (aSvc AuthSvc) createToken(userID uuid.UUID) (string, error) {
	secretKey := []byte(aSvc.secretKey)

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		userID.String(),
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
