package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ulexxander/meeting-time/db"
)

var ErrNoMeeting = errors.New("meeting does not exist")

type MeetingsService struct {
	queries *db.Queries
}

func NewMeetingsService(queries *db.Queries) *MeetingsService {
	return &MeetingsService{
		queries: queries,
	}
}

func (ms *MeetingsService) MeetingByID(ctx context.Context, id int) (*db.Meeting, error) {
	item, err := ms.queries.MeetingByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoMeeting
		}
	}
	return &item, nil
}

func (ms *MeetingsService) MeetingsBySchedule(ctx context.Context, scheduleID int) ([]db.Meeting, error) {
	return ms.queries.MeetingsBySchedule(ctx, scheduleID)
}

func (ms *MeetingsService) MeetingCreate(ctx context.Context, params db.MeetingCreateParams) (int, error) {
	return ms.queries.MeetingCreate(ctx, params)
}
