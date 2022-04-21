package storage

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrNoOrganization = errors.New("organization does not exist")

type OrganizationsStore struct {
	DB *sqlx.DB
}

type OrganizationInsertParams struct {
	Name string
}

const organizationGetByID = `SELECT * FROM organizations
WHERE id = $1`

func (os *OrganizationsStore) GetByID(id int) (*Organization, error) {
	var item Organization
	if err := os.DB.Get(&item, organizationGetByID, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoOrganization
		}
		return nil, err
	}
	return &item, nil
}

const organizationsCreate = `INSERT INTO organizations (name)
VALUES ($1)
RETURNING id`

func (os *OrganizationsStore) Create(params OrganizationInsertParams) (int, error) {
	var id int
	if err := os.DB.Get(&id, organizationsCreate, params.Name); err != nil {
		return 0, err
	}
	return id, nil
}
