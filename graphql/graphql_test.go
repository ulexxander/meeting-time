package graphql_test

import (
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
	"github.com/ulexxander/meeting-time/graphql"
	"github.com/ulexxander/meeting-time/graphql/client"
	"github.com/ulexxander/meeting-time/graphql/generated"
	"github.com/ulexxander/meeting-time/services"
	"github.com/ulexxander/meeting-time/testutil"
)

func TestGraphQL(t *testing.T) {
	ctx := testutil.Context(t)
	c := setupClient(t)

	var res struct {
		TeamByID struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"teamByID"`
	}
	query := `query NonexistentTeam ($id: ID!) {
		teamByID(id: $id) {
			id
			name
		}
	}`
	err := c.Query(ctx, query, client.Variables{"id": 123}, &res)
	require.NoError(t, err)
}

func setupClient(t *testing.T) *client.Client {
	queries := testutil.Queries(t)

	teamsService := services.NewTeamsService(queries)
	schedulesService := services.NewSchedulesService(queries)
	meetingsService := services.NewMeetingsService(queries)

	gqlResolver := graphql.NewResolver(
		teamsService,
		schedulesService,
		meetingsService,
	)
	gqlSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: gqlResolver,
	})
	gqlServer := handler.NewDefaultServer(gqlSchema)

	server := httptest.NewServer(gqlServer)
	t.Cleanup(server.Close)

	return &client.Client{
		URL: server.URL,
	}
}
