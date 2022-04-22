package graphql

import (
	"github.com/ulexxander/meeting-time/db"
	"github.com/ulexxander/meeting-time/graphql/model"
)

func convertSchedule(item *db.Schedule) *model.Schedule {
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

func convertSchedules(items []db.Schedule) []model.Schedule {
	converted := make([]model.Schedule, len(items))
	for i := range items {
		converted[i] = *convertSchedule(&items[i])
	}
	return converted
}
