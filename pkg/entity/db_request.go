package entity

type DBRequest struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	MinCount  int64  `json:"minCount,omitempty"`
	MaxCount  int64  `json:"maxCount,omitempty"`
}
