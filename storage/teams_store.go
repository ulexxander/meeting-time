package storage

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrNoTeam = errors.New("team does not exist")

type TeamsStore struct {
	DB *sqlx.DB
}

const teamGetByID = `SELECT * FROM teams
WHERE id = $1`

func (os *TeamsStore) GetByID(id int) (*Team, error) {
	var item Team
	if err := os.DB.Get(&item, teamGetByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoTeam
		}
		return nil, err
	}
	return &item, nil
}

const teamCreate = `INSERT INTO teams (name)
VALUES ($1)
RETURNING id`

type TeamCreateParams struct {
	Name string
}

func (os *TeamsStore) Create(params TeamCreateParams) (int, error) {
	var id int
	if err := os.DB.Get(&id, teamCreate, params.Name); err != nil {
		return 0, err
	}
	return id, nil
}
