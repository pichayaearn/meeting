package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/cmd/api/config"
	authRoute "github.com/pichayaearn/meeting/pkg/auth/route"
	authSvc "github.com/pichayaearn/meeting/pkg/auth/svc"
	"github.com/pichayaearn/meeting/pkg/middleware"
	"github.com/pichayaearn/meeting/pkg/route"

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
	meetingRepo := repo.NewMeetingRepo(db)
	commentRepo := repo.NewCommentRepo(db)

	userSvc := svc.NewUserSvc(svc.NewUserSvcCfgs{
		UserRepo: userRepo,
	})
	meetingSvc := svc.NewMeetingSvc(svc.NewMeetingSvcCfg{
		MeetingRepo: meetingRepo,
		UserRepo:    userRepo,
	})
	commentSvc := svc.NewCommentSvc(svc.NewCommentSvcCfg{
		CommentRepo: commentRepo,
		UserRepo:    userRepo,
	})
	authSvc := authSvc.NewAuthSvc(authSvc.NewAuthSvcCfg{
		UserSvc:   userSvc,
		SecretKey: cfg.SecretKey,
	})

	mw := middleware.Authenticate{
		Secret: cfg.SecretKey,
	}

	e.POST("/sign-up", authRoute.CreateUser(authRoute.CreateUserCfg{
		UserSvc: userSvc,
	}))

	e.POST("/login", authRoute.Login(authRoute.LoginCfg{
		AuthSvc: authSvc,
	}))

	e.POST("/meeting", route.CreateMeeting(route.CreateMeetingCfg{
		MeetingSvc: meetingSvc,
	}), mw.Authenticate)

	e.GET("/meetings", route.GetListMeeting(route.GetListMeetingCfg{
		MeetingSvc: meetingSvc,
	}), mw.Authenticate)

	e.GET("/comments", route.GetListComment(route.GetListCommentCfg{
		CommentSvc: commentSvc,
	}), mw.Authenticate)

	e.POST("/comment", route.CreateComment(route.CreateCommentCfg{
		CommentSvc: commentSvc,
	}), mw.Authenticate)

	return e

}
