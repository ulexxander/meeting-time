package services_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/db"
	"github.com/ulexxander/meeting-time/services"
)

func TestSchedulesService(t *testing.T) {
	ctx := testContext(t)
	queries := setupQueries(t)
	teamsService := services.NewTeamsService(queries)
	schedulesService := services.NewSchedulesService(queries)

	teamID, err := teamsService.TeamCreate(ctx, "Cool team")
	require.NoError(t, err)

	_, err = schedulesService.ScheduleByID(ctx, 123)
	require.ErrorIs(t, err, services.ErrNoSchedule)

	startsAt, _ := time.Parse(time.RFC3339, "2022-04-21 21:00:00+02:00")
	endsAt, _ := time.Parse(time.RFC3339, "2022-04-21 21:30:00+02:00")
	createdScheduleID, err := schedulesService.ScheduleCreate(ctx, db.ScheduleCreateParams{
		TeamID:   teamID,
		Name:     "Good schedule",
		StartsAt: startsAt,
		EndsAt:   endsAt,
	})
	require.NoError(t, err)

	scheduleByID, err := schedulesService.ScheduleByID(ctx, createdScheduleID)
	require.NoError(t, err)

	require.Equal(t, createdScheduleID, scheduleByID.ID)
	require.Equal(t, teamID, scheduleByID.TeamID)
	require.Equal(t, "Good schedule", scheduleByID.Name)
	require.Equal(t, startsAt, scheduleByID.StartsAt.UTC())
	require.Equal(t, endsAt, scheduleByID.EndsAt.UTC())
	require.NotZero(t, scheduleByID.CreatedAt)
	require.Nil(t, scheduleByID.UpdatedAt)

	teamSchedules, err := schedulesService.SchedulesByTeam(ctx, teamID)
	require.NoError(t, err)
	require.Len(t, teamSchedules, 1)
	require.Equal(t, *scheduleByID, teamSchedules[0])
}
