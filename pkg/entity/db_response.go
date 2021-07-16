package entity

import "time"

type DBResponse struct {
	startDate time.Time `json:"startDate,omitempty"`
	endDate   time.Time `json:"endDate,omitempty"`
	minCount  int64     `json:"minCount,omitempty"`
	maxCount  int64     `json:"maxCount,omitempty"`
}
