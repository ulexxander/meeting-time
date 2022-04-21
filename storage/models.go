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

type Schedule struct {
	ID        int
	TeamID    int
	Name      string
	StartsAt  time.Time
	EndsAt    time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time
}
