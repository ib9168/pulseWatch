package models

import "time"

type Service struct {
	ID                   int64
	UserID               int64
	Name                 string
	URL                  string
	CheckIntervalSeconds int //To check how often the worker can ping this services
	IsActive             bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
