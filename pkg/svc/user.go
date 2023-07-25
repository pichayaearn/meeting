package svc

import (
	"context"
	"errors"

	model "github.com/pichayaearn/meeting/pkg/model"
)

type UserSvc struct {
	userRepo model.UserRepo
}

type NewUserSvcCfgs struct {
	UserRepo model.UserRepo
}

func NewUserSvc(cfg NewUserSvcCfgs) model.UserSvc {
	return &UserSvc{
		userRepo: cfg.UserRepo,
	}
}

func (uSvc UserSvc) CreateUser(opts model.CreateUser) error {
	//check email exist
	userExist, err := uSvc.userRepo.Get(model.GetUserOpts{
		Email: opts.Email,
	}, context.Background())
	if err != nil {
		return err
	}

	if userExist != nil {
		//email already used
		return errors.New("email is already used")
	}

	//create user
	newUser, err := model.NewUser(opts.Email, opts.Password)
	if err != nil {
		return err
	}
	if err := uSvc.userRepo.Create(*newUser); err != nil {
		return err
	}

	return nil
}

func (uSvc UserSvc) GetUser(opts model.GetUserOpts, ctx context.Context) (*model.User, error) {
	return uSvc.userRepo.Get(opts, ctx)
}
