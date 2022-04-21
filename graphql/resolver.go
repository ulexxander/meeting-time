package graphql

import (
	"github.com/ulexxander/meeting-time/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TeamsStore storage.TeamsStore
}
