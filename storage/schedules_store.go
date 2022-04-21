package storage

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrNoSchedule = errors.New("schedule does not exist")

type SchedulesStore struct {
	DB *sqlx.DB
}

const scheduleGetByID = `SELECT * FROM schedules
WHERE id = $1`

func (ss *SchedulesStore) GetByID(id int) (*Schedule, error) {
	var item Schedule
	if err := ss.DB.Get(&item, scheduleGetByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoSchedule
		}
		return nil, err
	}
	return &item, nil
}

const schedulesGetByTeam = `SELECT * FROM schedules
WHERE teamID = $1`

func (ss *SchedulesStore) GetByTeam(teamID int) ([]Schedule, error) {
	var items []Schedule
	if err := ss.DB.Select(&items, schedulesGetByTeam, teamID); err != nil {
		return nil, err
	}
	return items, nil
}

const scheduleCreate = `INSERT INTO schedules (teamID, name, startsAt, endsAt)
VALUES ($1, $2, $3, $4)
RETURNING id`

type ScheduleCreateParams struct {
	TeamID   int
	Name     string
	StartsAt time.Time
	EndsAt   time.Time
}

func (ss *SchedulesStore) Create(params ScheduleCreateParams) (int, error) {
	var id int
	if err := ss.DB.Get(
		&id,
		scheduleCreate,
		params.TeamID,
		params.Name,
		params.StartsAt,
		params.EndsAt,
	); err != nil {
		return 0, err
	}
	return id, nil
}
