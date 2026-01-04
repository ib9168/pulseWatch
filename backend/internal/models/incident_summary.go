package models

import "time"

type IncidentSummary struct {
	ID         int64
	IncidentID int64
	CreatedAt  time.Time
	Model      string
	Prompt     string
	Summary    string
}
