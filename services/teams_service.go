package services

import "github.com/ulexxander/meeting-time/storage"

type TeamsService struct {
	teamsStore *storage.TeamsStore
}

func NewTeamsService(ts *storage.TeamsStore) *TeamsService {
	return &TeamsService{
		teamsStore: ts,
	}
}

func (ts *TeamsService) GetByID(id int) (*storage.Team, error) {
	return ts.teamsStore.GetByID(id)
}

func (ts *TeamsService) Create(params storage.TeamCreateParams) (int, error) {
	return ts.teamsStore.Create(params)
}
