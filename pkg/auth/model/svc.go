package model

type AuthSvc interface {
	Login(email, password string) (string, error)
}
