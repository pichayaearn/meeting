package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/cmd/api/config"
	authRoute "github.com/pichayaearn/meeting/pkg/auth/route"
	authSvc "github.com/pichayaearn/meeting/pkg/auth/svc"

	"github.com/pichayaearn/meeting/pkg/repo"
	"github.com/pichayaearn/meeting/pkg/svc"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun/extra/bundebug"
)

func newServer(cfg *config.Config) *echo.Echo {
	db := cfg.DB.MustNewDB()
	logger := logrus.New()
	logger.Info("new server")
	if cfg.Environment == "development" {
		db.AddQueryHook(bundebug.NewQueryHook())
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	e := echo.New()

	userRepo := repo.NewUserRepo(db)

	userSvc := svc.NewUserSvc(svc.NewUserSvcCfgs{
		UserRepo: userRepo,
	})
	authSvc := authSvc.NewAuthSvc(authSvc.NewAuthSvcCfg{
		UserSvc:   userSvc,
		SecretKey: cfg.SecretKey,
	})

	e.POST("/sign-up", authRoute.CreateUser(authRoute.CreateUserCfg{
		UserSvc: userSvc,
	}))

	e.POST("/login", authRoute.Login(authRoute.LoginCfg{
		AuthSvc: authSvc,
	}))

	return e

}
