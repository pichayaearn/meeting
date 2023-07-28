package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/uptrace/bun"
)

type meetingCommentBun struct {
	bun.BaseModel `bun:"table:meeting.comment_meetings"`
	ID            uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	MeetingID     uuid.UUID
	CommentID     uuid.UUID
	Status        string
	CreatedAt     time.Time
	CreatedBy     uuid.UUID
	UpdatedAt     time.Time
}

type commentBun struct {
	bun.BaseModel `bun:"table:meeting.comments"`
	ID            uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Detail        string
}

type CommentRepo struct {
	db *bun.DB
}

func NewCommentRepo(db *bun.DB) model.CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (cr CommentRepo) ListCommentID(opts model.GetListCommentOpts, ctx context.Context) ([]model.MeetingComment, error) {
	meetingComments := []meetingCommentBun{}
	q := cr.db.NewSelect().Model(&meetingComments)

	if opts.Limit > 0 {
		q.Limit(opts.Limit)
	}
	if opts.Offset > 0 {
		q.Offset(opts.Offset)
	}
	if err := q.OrderExpr("created_at DESC").ApplyQueryBuilder(addCommentFilter(opts)).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get list meeting comment error")
	}

	resp := []model.MeetingComment{}
	for _, v := range meetingComments {
		m, err := v.toModel()
		if err != nil {
			return nil, err
		}
		resp = append(resp, *m)
	}

	return resp, nil

}

func (cr CommentRepo) CreateCommentMeeting(meetingComment model.MeetingComment) error {
	mb := toMeetingCommentBun(meetingComment)
	if _, err := cr.db.NewInsert().Model(&mb).Exec(context.Background()); err != nil {
		return errors.New("create comment detail failed")
	}

	return nil

}

func (cr CommentRepo) CommentDetail(id uuid.UUID, ctx context.Context) (*model.CommentDetail, error) {
	comment := commentBun{}
	q := cr.db.NewSelect().Model(&comment)
	if err := q.Where("id = ?", id).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get comment detail error")
	}
	return comment.toModel()
}

func (cr CommentRepo) CreateCommentDetail(comment model.CommentDetail) (*model.CommentDetail, error) {
	cb := toCommentBun(comment)
	if _, err := cr.db.NewInsert().Model(&cb).Exec(context.Background()); err != nil {
		return nil, errors.New("create comment detail failed")
	}

	return cb.toModel()
}

func addCommentFilter(opts model.GetListCommentOpts) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.MeetingID != uuid.Nil {
			q.Where("meeting_id = ?", opts.MeetingID)
		}
		q.Where("status = ?", model.StatusCommentActive)

		return q

	}

}

func (mc meetingCommentBun) toModel() (*model.MeetingComment, error) {
	return model.MeetingCommentFactory(model.MeetingCommentFactoryOpts{
		ID:        mc.ID,
		MeetingID: mc.MeetingID,
		CommentID: mc.CommentID,
		Status:    mc.Status,
		CreatedBy: mc.CreatedBy,
		CreatedAt: mc.CreatedAt,
		UpdatedAt: mc.UpdatedAt,
	})
}

func (mc commentBun) toModel() (*model.CommentDetail, error) {
	return model.CommentDetailFactory(model.CommentDetailFactoryOpts{
		ID:     mc.ID,
		Detail: mc.Detail,
	})
}

func toCommentBun(comment model.CommentDetail) commentBun {
	return commentBun{
		ID:     comment.ID(),
		Detail: comment.Detail(),
	}
}

func toMeetingCommentBun(meetingComment model.MeetingComment) meetingCommentBun {
	return meetingCommentBun{
		ID:        meetingComment.ID(),
		MeetingID: meetingComment.MeetingID(),
		CommentID: meetingComment.CommentID(),
		Status:    string(meetingComment.Status()),
		CreatedAt: meetingComment.CreatedAt(),
		CreatedBy: meetingComment.CreatedByID(),
		UpdatedAt: meetingComment.UpdatedAt(),
	}
}
