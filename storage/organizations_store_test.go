package storage_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/storage"
)

func TestOrganizationsStore(t *testing.T) {
	db := setupDB(t)
	store := storage.OrganizationsStore{Store: &storage.Store{DB: db}}

	nonexistentOrg, err := store.GetByID(123)
	require.ErrorIs(t, err, storage.ErrNoOrganization)
	require.Nil(t, nonexistentOrg)

	createdOrg, err := store.Create(storage.OrganizationInsertParams{Name: "First!"})
	require.NoError(t, err)
	require.Equal(t, "First!", createdOrg.Name)

	orgByID, err := store.GetByID(createdOrg.ID)
	require.NoError(t, err)

	roundTimestamps(
		&createdOrg.CreatedAt,
		&createdOrg.UpdatedAt,
		&orgByID.CreatedAt,
		&orgByID.UpdatedAt,
	)

	require.Equal(t, createdOrg, orgByID)
}

func roundTimestamps(ts ...*time.Time) {
	for _, t := range ts {
		*t = t.Round(time.Millisecond)
	}
}
