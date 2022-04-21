package graphql

import (
	"github.com/ulexxander/meeting-time/graphql/model"
	"github.com/ulexxander/meeting-time/storage"
)

func convertSchedule(item *storage.Schedule) *model.Schedule {
	return &model.Schedule{
		ID:        item.ID,
		TeamID:    item.TeamID,
		Name:      item.Name,
		StartsAt:  item.StartsAt,
		EndsAt:    item.EndsAt,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func convertSchedules(items []storage.Schedule) []model.Schedule {
	converted := make([]model.Schedule, len(items))
	for i := range items {
		converted[i] = *convertSchedule(&items[i])
	}
	return converted
}
