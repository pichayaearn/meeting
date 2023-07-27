package svc

import (
	"context"
	"log"

	"github.com/pichayaearn/meeting/pkg/model"
)

type MeetingSvc struct {
	meetingRepo model.MeetingRepo
	userRepo    model.UserRepo
}

type NewMeetingSvcCfg struct {
	MeetingRepo model.MeetingRepo
	UserRepo    model.UserRepo
}

func NewMeetingSvc(cfg NewMeetingSvcCfg) model.MeetingSvc {
	return &MeetingSvc{
		meetingRepo: cfg.MeetingRepo,
		userRepo:    cfg.UserRepo,
	}
}

func (msvc MeetingSvc) Create(opts model.CreateMeetingOpts) error {
	newMeeting, err := model.NewMeeting(opts)
	if err != nil {
		return err
	}

	if err := msvc.meetingRepo.Create(*newMeeting); err != nil {
		return err
	}

	return nil

}

func (msvc MeetingSvc) List(opts model.GetMeetingOpts, ctx context.Context) ([]model.Meeting, error) {
	//get list meeting
	meetings, err := msvc.meetingRepo.List(opts, ctx)
	if err != nil {
		return nil, err
	}

	for i, v := range meetings {
		//find user detail
		log.Printf("createdBy %+v", v.CreatedByUUID())
		user, err := msvc.userRepo.Get(model.GetUserOpts{
			UserID: v.CreatedByUUID(),
		}, ctx)
		if err != nil {
			return nil, err
		}

		//set user created
		v.SetCreatedBy(*user)
		meetings[i] = v

	}

	return meetings, nil

}
