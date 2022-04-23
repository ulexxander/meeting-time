package services_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/db"
	"github.com/ulexxander/meeting-time/services"
)

func TestMeetingsService(t *testing.T) {
	ctx := testContext(t)
	queries := setupQueries(t)
	teamsService := services.NewTeamsService(queries)
	schedulesService := services.NewSchedulesService(queries)
	meetingsService := services.NewMeetingsService(queries)

	teamID, err := teamsService.TeamCreate(ctx, "My team")
	require.NoError(t, err)

	scheduleID, err := schedulesService.ScheduleCreate(ctx, db.ScheduleCreateParams{
		TeamID:   teamID,
		Name:     "Test schedule",
		StartsAt: time.Now(),
		EndsAt:   time.Now(),
	})
	require.NoError(t, err)

	_, err = meetingsService.MeetingByID(ctx, 123)
	require.ErrorIs(t, err, services.ErrNoMeeting)

	startedAt, _ := time.Parse(time.RFC3339, "2022-04-21 21:00:00+02:00")
	endedAt, _ := time.Parse(time.RFC3339, "2022-04-21 21:30:00+02:00")
	createdMeetingID, err := meetingsService.MeetingCreate(ctx, db.MeetingCreateParams{
		ScheduleID: scheduleID,
		StartedAt:  startedAt,
		EndedAt:    endedAt,
	})
	require.NoError(t, err)

	meetingByID, err := meetingsService.MeetingByID(ctx, createdMeetingID)
	require.NoError(t, err)

	require.Equal(t, createdMeetingID, meetingByID.ID)
	require.Equal(t, scheduleID, meetingByID.ScheduleID)
	require.Equal(t, startedAt, meetingByID.StartedAt.UTC())
	require.Equal(t, endedAt, meetingByID.EndedAt.UTC())
	require.NotZero(t, meetingByID.CreatedAt)
	require.Nil(t, meetingByID.UpdatedAt)
}
