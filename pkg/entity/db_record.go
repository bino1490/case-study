package entity

import "time"

type DBRecord struct {
	Key        string    `json:"key,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	TotalCount int64     `json:"totalCount"`
}
