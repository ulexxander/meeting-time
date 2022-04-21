package storage

import (
	"time"
)

type Organization struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
