package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/graphql/model"
	"github.com/ulexxander/meeting-time/storage"
)

func (r *mutationResolver) CreateOrganization(ctx context.Context, input model.OrganizationCreate) (*model.Organization, error) {
	item, err := r.OrganizationsStore.Create(storage.OrganizationInsertParams{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}
	return &model.Organization{
		ID:   int(item.ID),
		Name: item.Name,
	}, nil
}

func (r *queryResolver) OrganizationByID(ctx context.Context, id int) (*model.Organization, error) {
	item, err := r.OrganizationsStore.GetByID(uint(id))
	if err != nil {
		return nil, err
	}
	return &model.Organization{
		ID:   int(item.ID),
		Name: item.Name,
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
