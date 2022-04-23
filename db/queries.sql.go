// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: queries.sql

package db

import (
	"context"
	"time"
)

const meetingByID = `-- name: MeetingByID :one
SELECT id, schedule_id, started_at, ended_at, created_at, updated_at FROM meetings
WHERE id = $1
`

func (q *Queries) MeetingByID(ctx context.Context, id int) (Meeting, error) {
	row := q.db.QueryRowContext(ctx, meetingByID, id)
	var i Meeting
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.StartedAt,
		&i.EndedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const meetingCreate = `-- name: MeetingCreate :one
INSERT INTO meetings (schedule_id, started_at, ended_at)
VALUES ($1, $2, $3)
RETURNING id
`

type MeetingCreateParams struct {
	ScheduleID int
	StartedAt  time.Time
	EndedAt    time.Time
}

func (q *Queries) MeetingCreate(ctx context.Context, arg MeetingCreateParams) (int, error) {
	row := q.db.QueryRowContext(ctx, meetingCreate, arg.ScheduleID, arg.StartedAt, arg.EndedAt)
	var id int
	err := row.Scan(&id)
	return id, err
}

const meetingsBySchedule = `-- name: MeetingsBySchedule :many
SELECT id, schedule_id, started_at, ended_at, created_at, updated_at FROM meetings
WHERE schedule_id = $1
`

func (q *Queries) MeetingsBySchedule(ctx context.Context, scheduleID int) ([]Meeting, error) {
	rows, err := q.db.QueryContext(ctx, meetingsBySchedule, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Meeting
	for rows.Next() {
		var i Meeting
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.StartedAt,
			&i.EndedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const scheduleByID = `-- name: ScheduleByID :one
SELECT id, team_id, name, starts_at, ends_at, created_at, updated_at FROM schedules
WHERE id = $1
`

func (q *Queries) ScheduleByID(ctx context.Context, id int) (Schedule, error) {
	row := q.db.QueryRowContext(ctx, scheduleByID, id)
	var i Schedule
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Name,
		&i.StartsAt,
		&i.EndsAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const scheduleCreate = `-- name: ScheduleCreate :one
INSERT INTO schedules (team_id, name, starts_at, ends_at)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type ScheduleCreateParams struct {
	TeamID   int
	Name     string
	StartsAt time.Time
	EndsAt   time.Time
}

func (q *Queries) ScheduleCreate(ctx context.Context, arg ScheduleCreateParams) (int, error) {
	row := q.db.QueryRowContext(ctx, scheduleCreate,
		arg.TeamID,
		arg.Name,
		arg.StartsAt,
		arg.EndsAt,
	)
	var id int
	err := row.Scan(&id)
	return id, err
}

const schedulesByTeam = `-- name: SchedulesByTeam :many
SELECT id, team_id, name, starts_at, ends_at, created_at, updated_at FROM schedules
WHERE team_id = $1
`

func (q *Queries) SchedulesByTeam(ctx context.Context, teamID int) ([]Schedule, error) {
	rows, err := q.db.QueryContext(ctx, schedulesByTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Schedule
	for rows.Next() {
		var i Schedule
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.Name,
			&i.StartsAt,
			&i.EndsAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const teamByID = `-- name: TeamByID :one
SELECT id, name, created_at, updated_at FROM teams
WHERE id = $1
`

func (q *Queries) TeamByID(ctx context.Context, id int) (Team, error) {
	row := q.db.QueryRowContext(ctx, teamByID, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const teamCreate = `-- name: TeamCreate :one
INSERT INTO teams (name)
VALUES ($1)
RETURNING id
`

func (q *Queries) TeamCreate(ctx context.Context, name string) (int, error) {
	row := q.db.QueryRowContext(ctx, teamCreate, name)
	var id int
	err := row.Scan(&id)
	return id, err
}