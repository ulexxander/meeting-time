package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ulexxander/meeting-time/db"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/graphql/model"
)

func (r *mutationResolver) TeamCreate(ctx context.Context, input model.TeamCreate) (int, error) {
	id, err := r.TeamsService.TeamCreate(ctx, input.Name)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *mutationResolver) ScheduleCreate(ctx context.Context, input model.ScheduleCreate) (int, error) {
	id, err := r.SchedulesService.ScheduleCreate(ctx, db.ScheduleCreateParams{
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

func (r *queryResolver) TeamByID(ctx context.Context, id int) (*model.Team, error) {
	item, err := r.TeamsService.TeamByID(ctx, id)
	if err != nil {
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
	item, err := r.SchedulesService.ScheduleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertSchedule(item), nil
}

func (r *teamResolver) Schedules(ctx context.Context, obj *model.Team) ([]model.Schedule, error) {
	items, err := r.SchedulesService.SchedulesByTeam(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return convertSchedules(items), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
	teamResolver     struct{ *Resolver }
)
