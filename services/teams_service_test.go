package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/services"
)

func TestTeamsService(t *testing.T) {
	ctx := testContext(t)
	queries := setupQueries(t)
	teamsService := services.NewTeamsService(queries)

	_, err := teamsService.TeamByID(ctx, 123)
	require.ErrorIs(t, err, services.ErrNoTeam)

	createdTeamID, err := teamsService.TeamCreate(ctx, "Cool team")
	require.NoError(t, err)

	teamByID, err := teamsService.TeamByID(ctx, createdTeamID)
	require.NoError(t, err)

	require.Equal(t, createdTeamID, teamByID.ID)
	require.Equal(t, "Cool team", teamByID.Name)
	require.NotZero(t, teamByID.CreatedAt)
	require.Nil(t, teamByID.UpdatedAt)
}
