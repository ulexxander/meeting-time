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

func convertMeeting(item *db.Meeting) *model.Meeting {
	return &model.Meeting{
		ID:         item.ID,
		ScheduleID: item.ScheduleID,
		StartedAt:  item.StartedAt,
		EndedAt:    item.EndedAt,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
}

func convertMeetings(items []db.Meeting) []model.Meeting {
	converted := make([]model.Meeting, len(items))
	for i := range items {
		converted[i] = *convertMeeting(&items[i])
	}
	return converted
}
