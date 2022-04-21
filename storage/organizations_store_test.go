package storage_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/storage"
)

func TestOrganizationsStore(t *testing.T) {
	db := setupDB(t)
	store := storage.OrganizationsStore{DB: db}

	_, err := store.GetByID(123)
	require.ErrorIs(t, err, storage.ErrNoOrganization)

	createdOrgID, err := store.Create(storage.OrganizationInsertParams{Name: "First!"})
	require.NoError(t, err)

	orgByID, err := store.GetByID(createdOrgID)
	require.NoError(t, err)

	require.Equal(t, createdOrgID, orgByID.ID)
	require.Equal(t, "First!", orgByID.Name)
	require.NotZero(t, orgByID.CreatedAt)
	require.Nil(t, orgByID.UpdatedAt)
}
