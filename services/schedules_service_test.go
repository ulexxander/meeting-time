package services_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/services"
	"github.com/ulexxander/meeting-time/storage"
)

func TestSchedulesService(t *testing.T) {
	db := setupDB(t)
	teamsStore := &storage.TeamsStore{DB: db}
	schedulesStore := &storage.SchedulesStore{DB: db}
	teamsService := services.NewTeamsService(teamsStore)
	schedulesService := services.NewSchedulesService(schedulesStore)

	teamID, err := teamsService.Create(storage.TeamCreateParams{Name: "Cool team"})
	require.NoError(t, err)

	_, err = schedulesService.GetByID(123)
	require.ErrorIs(t, err, storage.ErrNoSchedule)

	startsAt, _ := time.Parse(time.RFC3339, "2022-04-21 21:00:00+02:00")
	endsAt, _ := time.Parse(time.RFC3339, "2022-04-21 21:30:00+02:00")
	createdScheduleID, err := schedulesService.Create(storage.ScheduleCreateParams{
		TeamID:   teamID,
		Name:     "Good schedule",
		StartsAt: startsAt,
		EndsAt:   endsAt,
	})
	require.NoError(t, err)

	scheduleByID, err := schedulesService.GetByID(createdScheduleID)
	require.NoError(t, err)

	require.Equal(t, createdScheduleID, scheduleByID.ID)
	require.Equal(t, teamID, scheduleByID.TeamID)
	require.Equal(t, "Good schedule", scheduleByID.Name)
	require.Equal(t, startsAt, scheduleByID.StartsAt.UTC())
	require.Equal(t, endsAt, scheduleByID.EndsAt.UTC())
	require.NotZero(t, scheduleByID.CreatedAt)
	require.Nil(t, scheduleByID.UpdatedAt)

	teamSchedules, err := schedulesService.GetByTeam(teamID)
	require.NoError(t, err)
	require.Len(t, teamSchedules, 1)
	require.Equal(t, *scheduleByID, teamSchedules[0])
}
