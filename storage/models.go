package storage

import (
	"time"
)

type Team struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
