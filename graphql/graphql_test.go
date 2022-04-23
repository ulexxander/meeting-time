package graphql_test

import (
	"flag"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/graphql"
	"github.com/ulexxander/meeting-time/graphql/client"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/graphql/model"
	"github.com/ulexxander/meeting-time/services"
	"github.com/ulexxander/meeting-time/testutil"
)

var flagQueryLog = flag.Bool("query-log", false, "Log GraphQL queries and responses")

func TestGraphQL(t *testing.T) {
	ctx := testutil.Context(t)
	c := setupClient(t)

	t.Run("nonexistent team", func(t *testing.T) {
		var res struct {
			TeamByID map[string]interface{} `json:"teamByID"`
		}
		query := `query ($id: ID!) {
			teamByID(id: $id) {
				id
			}
		}`
		err := c.Query(ctx, query, client.Variables{"id": 123}, &res)
		require.NoError(t, err)
		require.Nil(t, res.TeamByID)
	})

	t.Run("nonexistent schedule", func(t *testing.T) {
		var res struct {
			ScheduleByID map[string]interface{} `json:"scheduleByID"`
		}
		query := `query ($id: ID!) {
			scheduleByID(id: $id) {
				id
			}
		}`
		err := c.Query(ctx, query, client.Variables{"id": 123}, &res)
		require.NoError(t, err)
		require.Nil(t, res.ScheduleByID)
	})

	t.Run("nonexistent meeting", func(t *testing.T) {
		var res struct {
			MeetingByID map[string]interface{} `json:"meetingByID"`
		}
		query := `query ($id: ID!) {
			meetingByID(id: $id) {
				id
			}
		}`
		err := c.Query(ctx, query, client.Variables{"id": 123}, &res)
		require.NoError(t, err)
		require.Nil(t, res.MeetingByID)
	})

	var teamID int
	t.Run("creating team", func(t *testing.T) {
		var res struct {
			TeamCreate int `json:"teamCreate"`
		}
		query := `mutation ($input: TeamCreate!) {
			teamCreate(input: $input)
		}`
		input := model.TeamCreate{
			Name: "My team!",
		}
		err := c.Query(ctx, query, client.Variables{"input": input}, &res)
		require.NoError(t, err)
		teamID = res.TeamCreate
	})

	scheduleStartsAt, _ := time.Parse(time.Kitchen, "6:30AM")
	scheduleEndsAt, _ := time.Parse(time.Kitchen, "7:00AM")
	var scheduleID int
	t.Run("creating schedule", func(t *testing.T) {
		var res struct {
			ScheduleCreate int `json:"scheduleCreate"`
		}
		query := `mutation ($input: ScheduleCreate!) {
			scheduleCreate(input: $input)
		}`
		input := model.ScheduleCreate{
			TeamID:   teamID,
			Name:     "My schedule!",
			StartsAt: scheduleStartsAt,
			EndsAt:   scheduleEndsAt,
		}
		err := c.Query(ctx, query, client.Variables{"input": input}, &res)
		require.NoError(t, err)
		scheduleID = res.ScheduleCreate
	})

	meetingStartedAt := time.Now().UTC().Round(time.Microsecond)
	meetingEndedAt := meetingStartedAt.Add(time.Hour)
	var meetingID int
	t.Run("creating meeting", func(t *testing.T) {
		var res struct {
			MeetingCreate int `json:"meetingCreate"`
		}
		query := `mutation ($input: MeetingCreate!) {
			meetingCreate(input: $input)
		}`
		input := model.MeetingCreate{
			ScheduleID: scheduleID,
			StartedAt:  meetingStartedAt,
			EndedAt:    meetingEndedAt,
		}
		err := c.Query(ctx, query, client.Variables{"input": input}, &res)
		require.NoError(t, err)
		meetingID = res.MeetingCreate
	})

	t.Run("query team, schedules, meetings", func(t *testing.T) {
		var res struct {
			TeamByID struct {
				model.Team
				Schedules []struct {
					model.Schedule
					Meetings []model.Meeting `json:"meetings"`
				} `json:"schedules"`
			} `json:"teamByID"`
		}
		query := `query ($id: ID!) {
			teamByID(id: $id) {
				id
				name
				createdAt
				updatedAt
				schedules {
					id
					teamId
					name
					startsAt
					endsAt
					createdAt
					updatedAt
					meetings {
						id
						scheduleId
						startedAt
						endedAt
						createdAt
						updatedAt
					}
				}
			}
		}`
		err := c.Query(ctx, query, client.Variables{"id": teamID}, &res)
		require.NoError(t, err)

		require.Equal(t, teamID, res.TeamByID.ID)
		require.Equal(t, "My team!", res.TeamByID.Name)
		require.NotZero(t, res.TeamByID.CreatedAt)
		require.Nil(t, res.TeamByID.UpdatedAt)

		require.Len(t, res.TeamByID.Schedules, 1)
		schedule := res.TeamByID.Schedules[0]

		require.Equal(t, scheduleID, schedule.ID)
		require.Equal(t, teamID, schedule.TeamID)
		require.Equal(t, "My schedule!", schedule.Name)
		require.Equal(t, scheduleStartsAt, schedule.StartsAt)
		require.Equal(t, scheduleEndsAt, schedule.EndsAt)
		require.NotZero(t, schedule.CreatedAt)
		require.Nil(t, schedule.UpdatedAt)

		require.Len(t, schedule.Meetings, 1)
		meeting := schedule.Meetings[0]

		require.Equal(t, meetingID, meeting.ID)
		require.Equal(t, scheduleID, meeting.ScheduleID)
		require.Equal(t, meetingStartedAt, meeting.StartedAt)
		require.Equal(t, meetingEndedAt, meeting.EndedAt)
		require.NotZero(t, meeting.CreatedAt)
		require.Nil(t, meeting.UpdatedAt)
	})
}

func setupClient(t *testing.T) *client.Client {
	queries := testutil.Queries(t)

	teamsService := services.NewTeamsService(queries)
	schedulesService := services.NewSchedulesService(queries)
	meetingsService := services.NewMeetingsService(queries)

	gqlResolver := graphql.NewResolver(
		teamsService,
		schedulesService,
		meetingsService,
	)
	gqlSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: gqlResolver,
	})
	gqlServer := handler.NewDefaultServer(gqlSchema)

	server := httptest.NewServer(gqlServer)
	t.Cleanup(server.Close)

	c := &client.Client{
		URL: server.URL,
	}
	if *flagQueryLog {
		c.Logger = &client.LoggerStd{Logger: log.Default()}
	}

	return c
}
