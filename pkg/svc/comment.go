package svc

import (
	"context"

	"github.com/pichayaearn/meeting/pkg/model"
)

type CommentSvc struct {
	commentRepo model.CommentRepo
	userRepo    model.UserRepo
}

type NewCommentSvcCfg struct {
	CommentRepo model.CommentRepo
	UserRepo    model.UserRepo
}

func NewCommentSvc(cfg NewCommentSvcCfg) model.CommentSvc {
	return &CommentSvc{
		commentRepo: cfg.CommentRepo,
		userRepo:    cfg.UserRepo,
	}
}

func (csvc CommentSvc) List(opts model.GetListCommentOpts, ctx context.Context) ([]model.Comment, error) {
	resp := []model.Comment{}
	listMeetingComment, err := csvc.commentRepo.ListCommentID(opts, ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range listMeetingComment {
		//get comment by id
		commentDetail, err := csvc.commentRepo.CommentDetail(v.CommentID(), ctx)
		if err != nil {
			return nil, err
		}

		//find email created by
		findCreatedByEmail, err := csvc.userRepo.Get(model.GetUserOpts{
			UserID: v.CreatedByID(),
		}, ctx)
		if err != nil {
			return nil, err

		}
		if err := v.SetCreatedBy(findCreatedByEmail.Email()); err != nil {
			return nil, err
		}

		comment := model.Comment{}
		comment.SetMeetingComment(v)
		comment.SetCommentDetail(*commentDetail)

		resp = append(resp, comment)

	}

	return resp, nil
}

func (csvc CommentSvc) Create(opts model.CreateCommentOpts) error {
	meetingComment, commentDeatil, err := model.NewComment(opts)
	if err != nil {
		return err
	}

	//create comment detail
	cd, err := csvc.commentRepo.CreateCommentDetail(*commentDeatil)
	if err != nil {
		return err
	}
	if err := meetingComment.SetCommentID(cd.ID()); err != nil {
		return err
	}

	//create comment meeting
	if err := csvc.commentRepo.CreateCommentMeeting(*meetingComment); err != nil {
		return err
	}

	return nil

}
