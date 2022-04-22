package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ulexxander/meeting-time/db"
)

var ErrNoTeam = errors.New("team does not exist")

type TeamsService struct {
	queries *db.Queries
}

func NewTeamsService(queries *db.Queries) *TeamsService {
	return &TeamsService{
		queries: queries,
	}
}

func (ts *TeamsService) TeamByID(ctx context.Context, id int) (*db.Team, error) {
	item, err := ts.queries.TeamByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoTeam
		}
	}
	return &item, nil
}

func (ts *TeamsService) TeamCreate(ctx context.Context, name string) (int, error) {
	return ts.queries.TeamCreate(ctx, name)
}
