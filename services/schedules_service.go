package services

import "github.com/ulexxander/meeting-time/storage"

type SchedulesService struct {
	schedulesStore *storage.SchedulesStore
}

func NewSchedulesService(ss *storage.SchedulesStore) *SchedulesService {
	return &SchedulesService{
		schedulesStore: ss,
	}
}

func (ss *SchedulesService) GetByID(id int) (*storage.Schedule, error) {
	return ss.schedulesStore.GetByID(id)
}

func (ss *SchedulesService) GetByTeam(teamID int) ([]storage.Schedule, error) {
	return ss.schedulesStore.GetByTeam(teamID)
}

func (ss *SchedulesService) Create(params storage.ScheduleCreateParams) (int, error) {
	return ss.schedulesStore.Create(params)
}
