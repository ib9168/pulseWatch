package models

import "time"

type Incident struct {
	ID          int64
	ServiceID   int64
	StartedAt   time.Time
	ResolvedAt  *time.Time //so can point to nil if still open
	Status      string
	Title       string
	Description string
}
