package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/ulexxander/meeting-time/db"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/graphql/model"
	"github.com/ulexxander/meeting-time/services"
)

func (r *mutationResolver) TeamCreate(ctx context.Context, input model.TeamCreate) (int, error) {
	id, err := r.teamsService.TeamCreate(ctx, input.Name)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *mutationResolver) ScheduleCreate(ctx context.Context, input model.ScheduleCreate) (int, error) {
	id, err := r.schedulesService.ScheduleCreate(ctx, db.ScheduleCreateParams{
		TeamID:   input.TeamID,
		Name:     input.Name,
		StartsAt: input.StartsAt,
		EndsAt:   input.EndsAt,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *mutationResolver) MeetingCreate(ctx context.Context, input model.MeetingCreate) (int, error) {
	id, err := r.meetingsService.MeetingCreate(ctx, db.MeetingCreateParams{
		ScheduleID: input.ScheduleID,
		StartedAt:  input.StartedAt,
		EndedAt:    input.EndedAt,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *queryResolver) TeamByID(ctx context.Context, id int) (*model.Team, error) {
	item, err := r.teamsService.TeamByID(ctx, id)
	if err != nil {
		if errors.Is(err, services.ErrNoTeam) {
			return nil, nil
		}
		return nil, err
	}
	return &model.Team{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}, nil
}

func (r *queryResolver) ScheduleByID(ctx context.Context, id int) (*model.Schedule, error) {
	item, err := r.schedulesService.ScheduleByID(ctx, id)
	if err != nil {
		if errors.Is(err, services.ErrNoSchedule) {
			return nil, nil
		}
		return nil, err
	}
	return convertSchedule(item), nil
}

func (r *queryResolver) MeetingByID(ctx context.Context, id int) (*model.Meeting, error) {
	item, err := r.meetingsService.MeetingByID(ctx, id)
	if err != nil {
		if errors.Is(err, services.ErrNoMeeting) {
			return nil, nil
		}
		return nil, err
	}
	return convertMeeting(item), nil
}

func (r *scheduleResolver) Meetings(ctx context.Context, obj *model.Schedule) ([]model.Meeting, error) {
	items, err := r.meetingsService.MeetingsBySchedule(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return convertMeetings(items), nil
}

func (r *teamResolver) Schedules(ctx context.Context, obj *model.Team) ([]model.Schedule, error) {
	items, err := r.schedulesService.SchedulesByTeam(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return convertSchedules(items), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Schedule returns generated.ScheduleResolver implementation.
func (r *Resolver) Schedule() generated.ScheduleResolver { return &scheduleResolver{r} }

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
	scheduleResolver struct{ *Resolver }
	teamResolver     struct{ *Resolver }
)
