package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ulexxander/meeting-time/db"
)

var ErrNoSchedule = errors.New("schedule does not exist")

type SchedulesService struct {
	queries *db.Queries
}

func NewSchedulesService(queries *db.Queries) *SchedulesService {
	return &SchedulesService{
		queries: queries,
	}
}

func (ss *SchedulesService) ScheduleByID(ctx context.Context, id int) (*db.Schedule, error) {
	item, err := ss.queries.ScheduleByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoSchedule
		}
	}
	return &item, nil
}

func (ss *SchedulesService) SchedulesByTeam(ctx context.Context, teamID int) ([]db.Schedule, error) {
	return ss.queries.SchedulesByTeam(ctx, teamID)
}

func (ss *SchedulesService) ScheduleCreate(ctx context.Context, params db.ScheduleCreateParams) (int, error) {
	return ss.queries.ScheduleCreate(ctx, params)
}
