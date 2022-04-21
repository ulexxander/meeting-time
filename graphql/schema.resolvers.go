package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/graphql/model"
	"github.com/ulexxander/meeting-time/storage"
)

func (r *mutationResolver) TeamCreate(ctx context.Context, input model.TeamCreate) (int, error) {
	id, err := r.TeamsStore.Create(storage.TeamCreateParams{
		Name: input.Name,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *queryResolver) TeamByID(ctx context.Context, id int) (*model.Team, error) {
	item, err := r.TeamsStore.GetByID(id)
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
