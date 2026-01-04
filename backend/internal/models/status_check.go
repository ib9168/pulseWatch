package models

import "time"

type StatusCheck struct {
	ID             int64
	ServiceID      int64
	CheckedAt      time.Time
	status         string
	HTTPStatusCode *int
	ResponseTimeMS *int
	ErrorMessage   *string
}
