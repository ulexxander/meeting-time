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

	createdOrgID, err := service.Create(storage.TeamCreateParams{Name: "First!"})
	require.NoError(t, err)

	orgByID, err := service.GetByID(createdOrgID)
	require.NoError(t, err)

	require.Equal(t, createdOrgID, orgByID.ID)
	require.Equal(t, "First!", orgByID.Name)
	require.NotZero(t, orgByID.CreatedAt)
	require.Nil(t, orgByID.UpdatedAt)
}
