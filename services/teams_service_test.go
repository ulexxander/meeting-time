package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/services"
	"github.com/ulexxander/meeting-time/storage"
)

func TestTeamsService(t *testing.T) {
	db := setupDB(t)
	store := &storage.TeamsStore{DB: db}
	service := services.NewTeamsService(store)

	_, err := service.GetByID(123)
	require.ErrorIs(t, err, storage.ErrNoTeam)

	createdTeamID, err := service.Create(storage.TeamCreateParams{Name: "Cool team"})
	require.NoError(t, err)

	teamByID, err := service.GetByID(createdTeamID)
	require.NoError(t, err)

	require.Equal(t, createdTeamID, teamByID.ID)
	require.Equal(t, "Cool team", teamByID.Name)
	require.NotZero(t, teamByID.CreatedAt)
	require.Nil(t, teamByID.UpdatedAt)
}
