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

type meetingBun struct {
	bun.BaseModel `bun:"table:meeting.meetings"`
	ID            uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Title         string
	Detail        string
	Status        string
	CreatedAt     time.Time
	CreatedBy     uuid.UUID
	UpdatedAt     time.Time
}

type MeetingRepo struct {
	db *bun.DB
}

func NewMeetingRepo(db *bun.DB) model.MeetingRepo {
	return &MeetingRepo{
		db: db,
	}
}

func (mr MeetingRepo) Get(opts model.GetMeetingOpts, ctx context.Context) (*model.Meeting, error) {
	meeting := meetingBun{}
	q := mr.db.NewSelect().Model(&meeting)
	if err := q.OrderExpr("id DESC").ApplyQueryBuilder(addMeetingFilter(opts)).Limit(1).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get meeting error")
	}

	return meeting.toModel()
}

func (mr MeetingRepo) List(opts model.GetMeetingOpts, ctx context.Context) ([]model.Meeting, error) {
	meetings := []meetingBun{}
	q := mr.db.NewSelect().Model(&meetings)

	if opts.Limit > 0 {
		q.Limit(opts.Limit)
	}
	if opts.Offset > 0 {
		q.Offset(opts.Offset)
	}
	if err := q.OrderExpr("created_at DESC").ApplyQueryBuilder(addMeetingFilter(opts)).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("get list meeting error")
	}

	resp := []model.Meeting{}
	for _, v := range meetings {
		meeting, err := v.toModel()
		if err != nil {
			return nil, err
		}
		resp = append(resp, *meeting)
	}

	return resp, nil
}

func (mr MeetingRepo) Create(meeting model.Meeting) error {
	mb := toMeetingBun(meeting)
	if _, err := mr.db.NewInsert().Model(&mb).Exec(context.Background()); err != nil {
		return errors.New("create meeting failed")
	}
	return nil
}

func addMeetingFilter(opts model.GetMeetingOpts) func(q bun.QueryBuilder) bun.QueryBuilder {
	return func(q bun.QueryBuilder) bun.QueryBuilder {
		if opts.ID != uuid.Nil {
			q.Where("id = ?", opts.ID)
		}
		if opts.CreatedBy != uuid.Nil {
			q.Where("created_by = ?", opts.CreatedBy)
		}
		if opts.Status != "" {
			q.Where("status = ?", opts.Status)
		}

		return q

	}

}

func (mb meetingBun) toModel() (*model.Meeting, error) {
	return model.MeetingFactory(model.MeetingFactoryOpts{
		ID:        mb.ID,
		Title:     mb.Title,
		Detail:    mb.Detail,
		Status:    mb.Status,
		CreatedAt: mb.CreatedAt,
		CreatedBy: mb.CreatedBy,
		UpdatedAt: mb.UpdatedAt,
	})
}

func toMeetingBun(meeting model.Meeting) meetingBun {
	return meetingBun{
		ID:        meeting.ID(),
		Title:     meeting.Title(),
		Detail:    meeting.Detail(),
		Status:    string(meeting.Status()),
		CreatedAt: meeting.CreatedAt(),
		CreatedBy: meeting.CreatedByUUID(),
		UpdatedAt: meeting.UpdatedAt(),
	}
}
