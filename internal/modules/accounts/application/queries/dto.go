package queries

import "time"

type AccountDTO struct {
	ID        string
	UserID    string
	Label     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
