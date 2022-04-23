package graphql

import (
	"github.com/ulexxander/meeting-time/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	teamsService     *services.TeamsService
	schedulesService *services.SchedulesService
	meetingsService  *services.MeetingsService
}

func NewResolver(
	ts *services.TeamsService,
	ss *services.SchedulesService,
	ms *services.MeetingsService,
) *Resolver {
	return &Resolver{
		teamsService:     ts,
		schedulesService: ss,
		meetingsService:  ms,
	}
}
