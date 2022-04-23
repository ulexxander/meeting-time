
-- name: TeamByID :one
SELECT * FROM teams
WHERE id = $1;

-- name: TeamCreate :one
INSERT INTO teams (name)
VALUES ($1)
RETURNING id;

-- name: ScheduleByID :one
SELECT * FROM schedules
WHERE id = $1;

-- name: SchedulesByTeam :many
SELECT * FROM schedules
WHERE team_id = $1;

-- name: ScheduleCreate :one
INSERT INTO schedules (team_id, name, starts_at, ends_at)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: MeetingByID :one
SELECT * FROM meetings
WHERE id = $1;

-- name: MeetingsBySchedule :many
SELECT * FROM meetings
WHERE schedule_id = $1;

-- name: MeetingCreate :one
INSERT INTO meetings (schedule_id, started_at, ended_at)
VALUES ($1, $2, $3)
RETURNING id;
