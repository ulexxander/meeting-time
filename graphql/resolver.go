package graphql

import (
	"github.com/ulexxander/meeting-time/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TeamsService *services.TeamsService
}
